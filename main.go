package main

import (
	"SWP490_G21_Backend/controller"
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/ultity"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func init() {
	dbConfig := ultity.ReadDBConfig()
	stringConfig := dbConfig.DbUser + ":" +
		dbConfig.DbPassword + "@tcp(" +
		dbConfig.DbServer + ":" +
		dbConfig.DbPort + ")/" +
		dbConfig.Database + "?charset=utf8"
	orm.RegisterModel(
		new(model.Knowledge),
		new(model.Option),
		new(model.Question),
		new(model.User),
		new(model.ExamTest),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err1 := orm.RegisterDataBase("default", "mysql", stringConfig)
	if err1 != nil {
		fmt.Printf("false %v", err1)
	}

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	svConfig := ultity.ReadServerConfig()
	//start echo
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	//api
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.ListKnowledge)
	e.PUT("/knowledge", controller.KnowledgeUpload)
	e.GET("/history", controller.History, middleware.JWT([]byte("justAdmin")))
	e.GET("/history/:id", controller.GetExamById)
	e.GET("/api", controller.ApiWeb)
	e.POST("/login", controller.LoginResponse)
	e.POST("/register", controller.Register)
	e.PUT("/qa", controller.QaResponse, middleware.JWT([]byte("justAdmin")))
	//admin group
	admin := e.Group("/admin", middleware.JWT([]byte("justAdmin")))
	user := e.Group("/user", middleware.JWT([]byte("justUser")))
	user.GET("/history", controller.History)
	admin.GET("/user", controller.ListUser)
	/*
		request: adminToken
		response: list of user{id, username, role}
	*/
	admin.POST("/user", controller.AdminAddUser)
	/*
		request: adminToken, username, password, role
		response: id, username, role
	*/
	admin.GET("/user/:id", controller.GetUserById)
	admin.DELETE("/user/:id", controller.DeleteUserById)
	/*
		request: adminToken
		response: {"message":"delete user successfully"} or {"message":"delete user fail"}
	*/
	admin.PATCH("/user/:id", controller.UpdateUser)
	/*
		request: adminToken, role, change_password (true/...), password
		response: {"message":"edit user successfully"} or {"message":"edit user fail"}
	*/

	e.GET("/knowledge/:id", controller.DownloadKnowledge)
	/*
		request: adminToken
		response: file
	*/

	e.DELETE("/knowledge/:id", controller.DeleteKnowledge)
	/*
		request: adminToken
		response: {"message":"delete knowledge successfully"} or {"message":"delete knowledge fail"}
	*/

	e.GET("/test", func(context echo.Context) error {
		sess, _ := session.Get("session", context)
		sess.Options = &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7, // in seconds
		}
		dt := time.Now()
		sess.Values["username"] = dt.String()
		sess.Save(context.Request(), context.Response())
		return context.HTML(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":" + svConfig.PortBackend))
}
