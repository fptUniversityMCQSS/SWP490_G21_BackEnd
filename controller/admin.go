package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func ListUser(c echo.Context) error {
	o := orm.NewOrm()
	var user []*model.User
	var lists []*response.UserResponse

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userid := claims["userId"]
	fmt.Printf("%d abc \n", userid)
	//log.Println("test: " + userid.(string))

	qs, err := o.QueryTable("user").All(&user)

	//if has problem in connection
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}

	//add selected data to knowledge_Res list
	for _, u := range user {
		var us = new(response.UserResponse)
		us.Id = u.Id
		us.Username = u.Username
		us.Role = u.Role
		lists = append(lists, us)
	}
	fmt.Printf("%d ABC \n", qs)
	return c.JSON(http.StatusOK, lists)

}
