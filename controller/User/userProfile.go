package User

import (
	"SWP490_G21_Backend/model/entity"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ChangeProfile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	presentPassword := c.FormValue("password")
	newPassword := c.FormValue("newPassword")
	changePassword := c.FormValue("change_password")
	FullName := c.FormValue("fullName")
	Email := c.FormValue("email")
	Phone := c.FormValue("phone")

	IntUserId := int64(userId)
	user := &entity.User{
		Id: IntUserId,
	}

	if utility.CheckEmail(Email) {
		user.Email = Email
	} else {
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error065EmailInvalid,
		})
	}
	if utility.CheckPhone(Phone) {
		user.Phone = Phone
	} else {
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error066PhoneInvalid,
		})
	}
	if utility.CheckFullName(FullName) {
		user.FullName = FullName
	} else {
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error067FullNameInvalid,
		})
	}

	if changePassword == "true" {
		user.Username = userName
		user.Password = presentPassword
		err := utility.DB.Read(user, "username", "password")
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusBadRequest, response.Message{
				Message: utility.Error060CurrentPasswordInvalid,
			})
		}
		if utility.CheckPassword(newPassword) {
			user.Password = newPassword
			_, err := utility.DB.Update(user, "email", "phone", "full_name", "password")
			if err != nil {
				utility.FileLog.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error011UpdateUserFailed,
				})
			}
		} else {
			return c.JSON(http.StatusBadRequest, response.Message{
				Message: utility.Error064PasswordOfUserIsInvalid,
			})
		}

	} else {
		_, err := utility.DB.Update(user, "email", "phone", "full_name")
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error011UpdateUserFailed,
			})
		}
	}
	utility.FileLog.Println(userName + " changed Profile")
	return c.JSON(http.StatusOK, response.Message{
		Message: "Change Profile Successfully",
	})
}

func GetUserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	user := &entity.User{
		Username: userName,
	}
	err := utility.DB.Read(user, "username")
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error070NotFoundUserName,
		})
	}
	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
		Phone:    user.Phone,
		FullName: user.FullName,
	}
	utility.FileLog.Println(userName + " get user info")
	return c.JSON(http.StatusOK, userResponse)
}
