package controller

import (
	"SWP490_G21_Backend/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	RePassword := c.FormValue("re-password")

	user := &model.User{
		Username: Username,
	}
	o := orm.NewOrm()

	// Get a QuerySeter object. User is table name
	err := o.Read(user, "username")
	user.Role = "user"
	if Password != RePassword {
		return c.JSON(http.StatusBadRequest, "re enter password")
	}
	if err == nil {
		return c.JSON(http.StatusBadRequest, "user exist")
	}
	i, err := o.QueryTable("user").PrepareInsert()
	if err != nil {
		return err
	}
	user.Password = Password
	fmt.Println(i)
	insert, err := i.Insert(user)
	if err != nil {
		return err
	}
	fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return err1
	}

	//RePassword := c.FormValue("re-password")
	//var user []*model.User
	//o := orm.NewOrm()
	//// Get a QuerySeter object. User is table name
	//qs, err := o.QueryTable("user").All(&user)
	//
	////if has problem in connection
	//if err != nil {
	//	fmt.Println("querry error", err)
	//	return err
	//}
	//
	////Check if username is existed
	//for _, u := range user {
	//if Username == u.Username {
	//	return c.JSON(http.StatusBadRequest, "User da co")
	//	break
	//}else if Username != u.Username{
	//	if Password != RePassword{
	//		return c.JSON(http.StatusBadRequest, "Reenterpassword")
	//		break
	//	}else if Password == RePassword{
	//		user1 := &model.User{
	//			Username: Username,
	//			Password: Password,
	//		}
	//		i, err := o.QueryTable("user").PrepareInsert()
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(i)
	//		insert, err := i.Insert(user1)
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(insert)
	//		err1 := i.Close()
	//		if err1 != nil {
	//			return err1
	//		}
	//
	//	}
	//}}
	//fmt.Printf("%d register success \n", qs)
	return c.String(http.StatusOK, fmt.Sprintf("Register success "))
}
