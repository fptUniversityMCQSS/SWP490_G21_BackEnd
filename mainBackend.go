package main

import (
	"github.com/labstack/echo/v4"
	"lib/controller"
	"net/http"
)

type Test struct {
	abc string
	xxx string
	zzz int
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
		return context.JSON(http.StatusOK, Test{
			abc: "",
			xxx: "",
			zzz: 1233,
		})
	})
	//api
	e.Logger.Fatal(e.Start(":1323"))
}
