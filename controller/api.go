package controller

import "github.com/labstack/echo/v4"

func ApiWeb(c echo.Context) error {
	return c.File("views/apiWeb.html")
}
