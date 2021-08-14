package Authenticate

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"fmt"
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
	// Get a QuerySeter object. User is table name
	err := utility.DB.Read(user, "username")
	user.Role = "user"
	if err == nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error003UserExisted,
		})
	}
	i, err := utility.DB.QueryTable("user").PrepareInsert()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error004CantGetTableUser,
		})
	}
	user.Password = Password
	fmt.Println(i)
	insert, err := i.Insert(user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error005InsertUserError,
		})
	}
	fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}

	log.Printf(Username + " register success")
	return c.String(http.StatusOK, fmt.Sprintf("Register success "))
}
