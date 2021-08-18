package Admin

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ListUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	var user []*model.User
	var lists []*response.UserResponse

	_, err := utility.DB.QueryTable("user").All(&user)

	//if has problem in connection
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error002ErrorQueryForGetAllUsers,
		})
	}

	//add selected data to knowledge_Res list
	for _, u := range user {
		var us = new(response.UserResponse)
		us.Id = u.Id
		us.Username = u.Username
		us.Role = u.Role
		us.Email = u.Email
		us.Phone = u.Phone
		us.FullName = u.FullName
		lists = append(lists, us)
	}
	utility.FileLog.Println(userName + " get list user ")
	return c.JSON(http.StatusOK, lists)

}

func AddUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	FullName := c.FormValue("fullName")
	Email := c.FormValue("email")
	Phone := c.FormValue("phone")
	role := c.FormValue("role")

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	user := &model.User{
		Username: username,
	}

	if utility.CheckUsername(username) {
		err := utility.DB.Read(user, "username")

		if err == nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error003UserExisted,
			})
		}
		i, err := utility.DB.QueryTable("user").PrepareInsert()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error004CantGetTableUser,
			})
		}
		if utility.CheckRole(role) {
			user.Role = role
		} else {
			utility.FileLog.Println("Role of user invalid")
			return c.JSON(http.StatusBadRequest, response.Message{
				Message: utility.Error063RoleOfUserIsInvalid,
			})
		}
		if utility.CheckPassword(password) {
			user.Password = password
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
		insert, err := i.Insert(user)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error005InsertUserError,
			})
		}
		err1 := i.Close()
		if err1 != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error022CloseConnectionError,
			})
		}
		userResponse := response.UserResponse{
			Id:       insert,
			Username: username,
			FullName: FullName,
			Role:     role,
			Email:    Email,
			Phone:    Phone,
		}
		utility.FileLog.Println(userName + " add new user has id: " + strconv.Itoa(int(insert)))
		return c.JSON(http.StatusOK, userResponse)
	} else {
		utility.FileLog.Println("Username must not contains special characters and has length at least 8 characters")
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error006UserNameModified,
		})
	}
}

func GetUserById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error008UserIdInvalid,
		})
	}
	var user model.User

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	err = utility.DB.QueryTable("user").Filter("id", id).One(&user)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error007CantGetUser,
		})
	}

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		FullName: user.FullName,
		Role:     user.Role,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	utility.FileLog.Println(userName + " get user " + user.Username)
	return c.JSON(http.StatusOK, userResponse)
}

func DeleteUserById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error008UserIdInvalid,
		})
	}
	userExited := utility.DB.QueryTable("user").Filter("id", id).Exist()
	if userExited == false {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error068UserDoesNotExist,
		})
	}
	_, err = utility.DB.QueryTable("user").Filter("id", id).Delete()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error009DeleteUserFailed,
		})
	}
	return c.JSON(http.StatusOK, response.Message{
		Message: "Delete User successfully",
	})
}

func UpdateUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	changePassword := c.FormValue("change_password")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error008UserIdInvalid,
		})
	}
	user := &model.User{
		Id: id,
	}

	role := c.FormValue("role")
	if utility.CheckRole(role) {
		user.Role = role
	} else {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error010RoleOfUserIsInvalid,
		})
	}
	if changePassword == "true" {
		password := c.FormValue("password")
		if utility.CheckPassword(password) {
			user.Password = password
			_, err := utility.DB.Update(user, "role", "password")
			if err != nil {
				utility.FileLog.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error011UpdateUserFailed,
				})
			}
		} else {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusBadRequest, response.Message{
				Message: utility.Error064PasswordOfUserIsInvalid,
			})
		}
	} else {
		_, err := utility.DB.Update(user, "role")
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error011UpdateUserFailed,
			})
		}
	}
	utility.FileLog.Println(userName + "edit user: " + user.Username)
	return c.JSON(http.StatusOK, response.Message{
		Message: "edit user successfully",
	})
}
