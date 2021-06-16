package controller

import "github.com/labstack/echo/v4"

func Qa(c echo.Context) error {
	return c.File("views/qa.html")
}
