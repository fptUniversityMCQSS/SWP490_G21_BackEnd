package User

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ChangePassword(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	presentPassword := c.FormValue("password")
	newPassword := c.FormValue("newPassword")
	user := &model.User{
		Username: userName,
		Password: presentPassword,
	}
	// Get a QuerySeter object. User is table name
	err := utility.DB.Read(user, "username", "password")
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error060CurrentPasswordInvalid,
		})
	}
	user.Password = newPassword
	_, err = utility.DB.Update(user)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error011UpdateUserFailed,
		})
	}

	utility.FileLog.Println(userName + " changed password")
	return c.JSON(http.StatusOK, response.Message{
		Message: "Change Password Successfully",
	})
}
