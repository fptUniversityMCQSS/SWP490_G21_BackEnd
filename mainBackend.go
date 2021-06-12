package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"lib/controller"
	"lib/model"
	"net/http"
)

func init() {
	orm.RegisterModel(new(model.Knowledge), new(model.Option), new(model.Question), new(model.User))
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err1 := orm.RegisterDataBase("default", "mysql", "root:abc@/question_answer_db?charset=utf8")
	if err1 != nil {
		fmt.Printf("false %v", err1)
	}
}
func main() {
	//start echo
	e := echo.New()
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.Knowledge)
	e.GET("/knowledge/upload", controller.KnowledgeUpload)
	e.GET("/history", controller.History)
	e.GET("/api", controller.ApiWeb)

	e.GET("/test", func(context echo.Context) error {
		return context.JSON(http.StatusOK, []model.Question{})
	})
	//api
	e.Logger.Fatal(e.Start(":1323"))
}
