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
	fmt.Printf("%d \n", userid)
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
	fmt.Printf("", qs)
	return c.JSON(http.StatusOK, lists)

}

func AdminAddUser(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	Role := c.FormValue("role")

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userid := claims["userId"]
	fmt.Printf("%d \n", userid)

	user := &model.User{
		Username: Username,
	}
	o := orm.NewOrm()

	// Get a QuerySeter object. User is table name
	err := o.Read(user, "username")

	if err == nil {
		return c.JSON(http.StatusBadRequest, "user exist")
	}
	i, err := o.QueryTable("user").PrepareInsert()
	if err != nil {
		return err
	}
	user.Password = Password
	user.Role = Role

	insert, err := i.Insert(user)
	if err != nil {
		return err
	}
	fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return err1
	}

	return c.String(http.StatusOK, fmt.Sprintf("Add user success "))
}