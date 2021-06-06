package main

import (
	"github.com/labstack/echo/v4"
	"lib/controller"
)

func main() {
	e := echo.New()
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.Knowledge)
	e.GET("/knowledge/upload", controller.KnowledgeUpload)
	e.GET("/history", controller.History)
	e.GET("/api", controller.ApiWeb)
	//api
	e.Logger.Fatal(e.Start(":1323"))
}
