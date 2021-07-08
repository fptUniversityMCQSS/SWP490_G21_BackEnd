package controller

import (
	"SWP490_G21_Backend/model"
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
		return c.JSON(http.StatusInternalServerError, "Delete User Failed")
	}
	return c.JSON(http.StatusOK, "Delete User successfully")
}

func UpdateUser(c echo.Context) error {
	changePassword := c.FormValue("change_password")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	o := orm.NewOrm()
	user := &model.User{
		Id: id,
	}
	if changePassword == "true" {
		role := c.FormValue("role")
		password := c.FormValue("password")
		if role != "" {
			user.Role = role
			_, err := o.Update(user, "role")
			if err != nil {
				return err
			}
		}
		if password != "" {
			user.Password = password
			_, err := o.Update(user, "password")
			if err != nil {
				return err
			}
		}
		return c.JSON(http.StatusOK, "edit user successfully")
	}
	return c.JSON(http.StatusInternalServerError, "edit user fail")
}
