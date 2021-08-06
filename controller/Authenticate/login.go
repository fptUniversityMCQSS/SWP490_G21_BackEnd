package Authenticate

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

const (
	JwtSignature = "SWP490_G21"
)

func LoginResponse(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	user := &model.User{
		Username: Username,
		Password: Password,
	}

	// Get a QuerySeter object. User is table name
	err := utility.DB.Read(user, "username", "password")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "InvalidUser")
	}

	token := jwt.New(jwt.SigningMethodHS256) //header

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["userId"] = user.Id
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(120 * time.Minute).Unix() // payload

	t, err := token.SignedString([]byte(JwtSignature))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "login fail")
	}
	log.Printf(Username + "login success")
	return c.JSON(http.StatusOK, &response.LoginResponse{Username: user.Username, Role: user.Role, Token: t})

}
