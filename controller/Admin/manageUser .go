package Admin

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/ultity"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ListUser(c echo.Context) error {
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	o := orm.NewOrm()
	var user []*model.User
	var lists []*response.UserResponse

	_, err := o.QueryTable("user").All(&user)

	//if has problem in connection
	if err != nil {
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
	log.Printf(username + " get list user ")
	return c.JSON(http.StatusOK, lists)

}

func AdminAddUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	role := c.FormValue("role")

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	//userid := claims["userId"]
	Username := claims["username"].(string)
	//StringId := strconv.FormatInt(int64(userid), 64)

	user := &model.User{
		Username: username,
	}
	o := orm.NewOrm()
	if ultity.CheckUsername(username) {
		err := o.Read(user, "username")

		if err == nil {
			return c.JSON(http.StatusBadRequest, "user exist")
		}
		i, err := o.QueryTable("user").PrepareInsert()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "query error",
			})
		}
		if ultity.CheckRole(role) {
			user.Role = role
		}
		if ultity.CheckPassword(password) {
			user.Password = password
		}

		insert, err := i.Insert(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Insert error",
			})
		}
		err1 := i.Close()
		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Close connection error",
			})
		}
		userResponse := response.UserResponse{
			Id:       insert,
			Username: username,
			Role:     role,
		}
		log.Printf(Username + " add new user has id: ")
		return c.JSON(http.StatusOK, userResponse)
	} else {
		log.Printf("Username exist")
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Username exist",
		})
	}
}

func GetUserById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	o := orm.NewOrm()
	var user model.User

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	//userid := claims["userId"]
	Username := claims["username"].(string)

	err := o.QueryTable("user").Filter("id", id).One(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}
	log.Printf(Username + " get user " + user.Username)
	return c.JSON(http.StatusOK, userResponse)
}

func DeleteUserById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := &model.User{
		Id: id,
	}
	o := orm.NewOrm()
	_, err := o.Delete(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Delete User Failed",
		})
	}
	return c.JSON(http.StatusOK, response.Message{
		Message: "Delete User successfully",
	})
}

func UpdateUser(c echo.Context) error {
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	//userid := claims["userId"]
	Username := claims["username"].(string)

	changePassword := c.FormValue("change_password")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	o := orm.NewOrm()
	user := &model.User{
		Id: id,
	}
	role := c.FormValue("role")
	if ultity.CheckRole(role) {
		user.Role = role
	} else {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "edit user failed",
		})
	}
	if changePassword == "true" {
		password := c.FormValue("password")
		if ultity.CheckPassword(password) {
			user.Password = password
			_, err := o.Update(user, "role", "password")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: "edit user failed",
				})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "edit user failed",
			})
		}
	} else {
		_, err := o.Update(user, "role")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "edit user failed",
			})
		}
	}
	log.Printf(Username + "edit user: " + user.Username)
	return c.JSON(http.StatusOK, response.Message{
		Message: "edit user successfully",
	})
}
