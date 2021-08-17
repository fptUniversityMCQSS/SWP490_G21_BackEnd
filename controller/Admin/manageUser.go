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
		lists = append(lists, us)
	}
	utility.FileLog.Println(userName + " get list user ")
	return c.JSON(http.StatusOK, lists)

}

func AddUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
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
		}
		if utility.CheckPassword(password) {
			user.Password = password
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
			Role:     role,
		}
		utility.FileLog.Println(userName + " add new user has id: " + strconv.Itoa(int(insert)))
		return c.JSON(http.StatusOK, userResponse)
	} else {
		utility.FileLog.Println("Username empty")
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error006UserNameEmpty,
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
		Role:     user.Role,
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
	user := &model.User{
		Id: id,
	}
	_, err = utility.DB.Delete(user)
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
		return c.JSON(http.StatusInternalServerError, response.Message{
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
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error012PasswordEmpty,
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
