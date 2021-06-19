package main

import (
	"github.com/labstack/echo/v4"
	"lib/controller"
)


func main() {
	//start echo
	e := echo.New()
	//website
	e.GET("/", controller.Home)
	e.GET("/home", controller.Home)
	e.GET("/qa", controller.Qa)
	e.GET("/knowledge", controller.Knowledge)
	e.GET("/knowledge/upload", controller.KnowledgeUpload)
	e.POST("/knowledge/upload", controller.KnowledgeUpload)
	e.GET("/history", controller.History)
	e.GET("/api", controller.ApiWeb)

	//e.GET("/test", func(context echo.Context) error {
	//	return context.JSON(http.StatusOK, []model.Question{
	//		{
	//			Content: "abc1",
	//		},
	//		{
	//			Content: "abc2",
	//		},
	//		{
	//			Content: "abc3",
	//		},
	//	})
	//})
	//api
	e.Logger.Fatal(e.Start(":1323"))
}
