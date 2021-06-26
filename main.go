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
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.ListKnowledge)
	e.PUT("/knowledge", controller.KnowledgeUpload)
	e.GET("/history", controller.History, middleware.JWT([]byte("justAdmin")))
	e.GET("/api", controller.ApiWeb)
	e.POST("/login", controller.LoginResponse)
	e.PUT("/qa", controller.QaResponse, middleware.JWT([]byte("justAdmin")))
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
	//api

	e.Logger.Fatal(e.Start(":" + svConfig.PortBackend))
}
