package User

import (
	"SWP490_G21_Backend/model"
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
	presentPassword := c.FormValue("password")
	newPassword := c.FormValue("newPassword")
	FullName := c.FormValue("fullName")
	Email := c.FormValue("email")
	Phone := c.FormValue("phone")
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
	if utility.CheckPassword(newPassword) {
		user.Password = newPassword
	} else {
		utility.FileLog.Println("Password has at least 8 character")
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error064PasswordOfUserIsInvalid,
		})
	}
	if utility.CheckEmail(Email) {
		user.Email = Email
	} else {
		utility.FileLog.Println("Email has a form xxx@xxx.xxx")
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error065EmailInvalid,
		})
	}
	if utility.CheckPhone(Phone) {
		user.Phone = Phone
	} else {
		utility.FileLog.Println("Phone must be 10 digit")
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error066PhoneInvalid,
		})
	}
	if utility.CheckFullName(FullName) {
		user.FullName = FullName
	} else {
		utility.FileLog.Println("Full Name has 8 to 30 characters")
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error067FullNameInvalid,
		})
	}
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

func GetUserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	user := &model.User{
		Username: userName,
	}
	err := utility.DB.Read(user, "username")
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error001InvalidUser,
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
