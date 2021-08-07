package Admin

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
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
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
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
	log.Printf(userName + " get list user ")
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
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "user existed",
			})
		}
		i, err := utility.DB.QueryTable("user").PrepareInsert()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "dont have table user",
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
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Insert user error",
			})
		}
		err1 := i.Close()
		if err1 != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Close connection error",
			})
		}
		userResponse := response.UserResponse{
			Id:       insert,
			Username: username,
			Role:     role,
		}
		log.Printf(userName + " add new user has id: ")
		return c.JSON(http.StatusOK, userResponse)
	} else {
		log.Printf("Username empty")
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Username empty",
		})
	}
}

func GetUserById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var user model.User

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	err := utility.DB.QueryTable("user").Filter("id", id).One(&user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "dont have user",
		})
	}

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}
	log.Printf(userName + " get user " + user.Username)
	return c.JSON(http.StatusOK, userResponse)
}

func DeleteUserById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "user id invalid",
		})
	}
	user := &model.User{
		Id: id,
	}
	_, err = utility.DB.Delete(user)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Delete User Failed",
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
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "user id invalid",
		})
	}
	user := &model.User{
		Id: id,
	}
	role := c.FormValue("role")
	if utility.CheckRole(role) {
		user.Role = role
	} else {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "edit user failed",
		})
	}
	if changePassword == "true" {
		password := c.FormValue("password")
		if utility.CheckPassword(password) {
			user.Password = password
			_, err := utility.DB.Update(user, "role", "password")
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: "edit user failed",
				})
			}
		} else {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "edit user failed",
			})
		}
	} else {
		_, err := utility.DB.Update(user, "role")
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "edit user failed",
			})
		}
	}
	log.Printf(userName + "edit user: " + user.Username)
	return c.JSON(http.StatusOK, response.Message{
		Message: "edit user successfully",
	})
}
