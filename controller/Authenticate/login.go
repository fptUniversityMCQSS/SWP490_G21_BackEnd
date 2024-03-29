package Authenticate

import (
	"SWP490_G21_Backend/model/entity"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	JwtSignature = "SWP490_G21"
)

func LoginResponse(c echo.Context) error {

	Username := c.FormValue("username")
	Password := c.FormValue("password")
	user := &entity.User{
		Username: Username,
	}
	err := utility.DB.Read(user, "username")
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error070NotFoundUserName,
		})
	}
	// Get a QuerySeter object. User is table name

	if user.Password != Password {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error001InvalidUser,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256) //header

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["userId"] = user.Id
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(1440 * time.Minute).Unix() // payload

	t, err := token.SignedString([]byte(JwtSignature))
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error013CreateTokenOfUserFailed,
		})
	}
	utility.FileLog.Println(Username + " login success")
	return c.JSON(http.StatusOK, &response.LoginResponse{Username: user.Username, FullName: user.FullName, Role: user.Role, Email: user.Email, Phone: user.Phone, Token: t})

}
