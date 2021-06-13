package controller

import (
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"lib/model"
	"net/http"
	"time"
)

func LoginResponse(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	user := &model.User{
		Username: Username,
		Password: Password,
	}
	o := orm.NewOrm()

	// Get a QuerySeter object. User is table name
	err := o.Read(user, "username")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "InvalidUser")
	}

	token := jwt.New(jwt.SigningMethodHS256) //header

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(3 * time.Minute).Unix() // payload
	if user.Role == "user" {
		t, err := token.SignedString([]byte("justUser"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &model.LoginResponse{Token: t})
	} else if user.Role == "admin" {
		t, err := token.SignedString([]byte("justAdmin"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &model.LoginResponse{Token: t})

	}
	//signature

	return nil

}
