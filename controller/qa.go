package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
)

func Qa(c echo.Context) error {
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, string(content))
}
