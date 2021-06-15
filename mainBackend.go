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
	svConfig := ultity.ReadDBConfig()
	stringConfig := svConfig.DbUser + ":" + svConfig.DbPassword + "@/" + svConfig.Database + "?charset=utf8"
	orm.RegisterModel(new(model.Knowledge), new(model.Option), new(model.Question), new(model.User))
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err1 := orm.RegisterDataBase("default", "mysql", stringConfig)
	if err1 != nil {
		fmt.Printf("false %v", err1)
	}

}
func main() {
	//start echo
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.POST("/qa", controller.QaResponse)
	e.GET("/knowledge", controller.Knowledge)
	e.GET("/knowledge/upload", controller.KnowledgeUpload)
	e.GET("/history", controller.History)
	e.GET("/api", controller.ApiWeb)
	e.POST("/login", controller.LoginResponse)

	e.GET("/test", func(context echo.Context) error {
		return context.JSON(http.StatusOK, []model.Question{})
	})
	//api

	e.Logger.Fatal(e.Start(":1323"))
}
