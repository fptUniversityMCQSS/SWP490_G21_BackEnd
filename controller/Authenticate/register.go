package Authenticate

import (
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/model/unity"
	"SWP490_G21_Backend/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Register(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	FullName := c.FormValue("fullName")
	Email := c.FormValue("email")
	Phone := c.FormValue("phone")
	user := &unity.User{
		Username: Username,
	}
	// Get a QuerySeter object. User is table name
	if !utility.CheckPassword(Username) {
		utility.FileLog.Println(utility.Error006UserNameModified)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error006UserNameModified,
		})
	}
	err := utility.DB.Read(user, "username")
	user.Role = "user"
	if err == nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error003UserExisted,
		})
	}
	i, err := utility.DB.QueryTable("user").PrepareInsert()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error004CantGetTableUser,
		})
	}
	if utility.CheckPassword(Password) {
		user.Password = Password
	} else {
		utility.FileLog.Println(utility.Error064PasswordOfUserIsInvalid)
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
	insert, err := i.Insert(user)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error005InsertUserError,
		})
	}
	err1 := i.Close()
	if err1 != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}
	utility.FileLog.Println(Username + " register success with id: " + strconv.Itoa(int(insert)))
	var userResponse = response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
		Phone:    user.Phone,
		FullName: user.FullName,
	}
	return c.JSON(http.StatusOK, userResponse)
}
