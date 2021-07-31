package Authenticate

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Register(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")

	user := &model.User{
		Username: Username,
	}
	o := orm.NewOrm()

	// Get a QuerySeter object. User is table name
	err := o.Read(user, "username")
	user.Role = "user"
	if err == nil {
		return c.JSON(http.StatusBadRequest, "user exist")
	}
	i, err := o.QueryTable("user").PrepareInsert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}
	user.Password = Password
	fmt.Println(i)
	insert, err := i.Insert(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Insert error",
		})
	}
	fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close error",
		})
	}

	log.Printf(Username + " register success")
	return c.String(http.StatusOK, fmt.Sprintf("Register success "))
}
