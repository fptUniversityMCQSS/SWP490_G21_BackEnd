package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"lib/controller"
	"lib/model"
	"lib/ultity"
	"net/http"
)

func init() {
	dbConfig := ultity.ReadDBConfig()
	stringConfig := dbConfig.DbUser + ":" +
		dbConfig.DbPassword + "@tcp(" +
		dbConfig.DbServer + ":" +
		dbConfig.DbPort + ")/" +
		dbConfig.Database + "?charset=utf8"
	orm.RegisterModel(new(model.Knowledge), new(model.Option), new(model.Question), new(model.User))
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
		AllowMethods: []string{"GET", "POST", "PUT"},
	}))
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.Knowledge)
	e.PUT("/knowledge", controller.KnowledgeUpload)
	e.GET("/history", controller.History)
	e.GET("/api", controller.ApiWeb)
	e.POST("/login", controller.LoginResponse)
	e.PUT("/qa", controller.QaResponse, middleware.JWT([]byte("justAdmin")))

	e.GET("/test", func(context echo.Context) error {
		return context.JSON(http.StatusOK, []model.Question{})
	})
	//api

	e.Logger.Fatal(e.Start(":" + svConfig.PortBackend))
}
