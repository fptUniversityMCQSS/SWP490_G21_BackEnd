package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/ultity"
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

	_, err := o.QueryTable("user").All(&user)

	//if has problem in connection
	if err != nil {
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
	return c.JSON(http.StatusOK, lists)

}

func AdminAddUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	role := c.FormValue("role")

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userid := claims["userId"]
	fmt.Printf("%d \n", userid)

	user := &model.User{
		Username: username,
	}
	o := orm.NewOrm()
	if ultity.CheckUsername(username) {
		err := o.Read(user, "username")

		if err == nil {
			return c.JSON(http.StatusBadRequest, "user exist")
		}
		i, err := o.QueryTable("user").PrepareInsert()
		if err != nil {
			return err
		}
		if ultity.CheckRole(role) {
			user.Role = role
		}
		if ultity.CheckPassword(password) {
			user.Password = password
		}

		insert, err := i.Insert(user)
		if err != nil {
			return err
		}
		fmt.Println(insert)
		err1 := i.Close()
		if err1 != nil {
			return err1
		}
		userResponse := response.UserResponse{
			Id:       insert,
			Username: username,
			Role:     role,
		}
		return c.JSON(http.StatusOK, userResponse)
	} else {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Username exist",
		})
	}
}
