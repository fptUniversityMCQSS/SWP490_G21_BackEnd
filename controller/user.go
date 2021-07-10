package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/ultity"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

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
				return err
			}
		} else {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "edit user failed",
			})
		}
	} else {
		_, err := o.Update(user, "role")
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, response.Message{
		Message: "edit user successfully",
	})
}
