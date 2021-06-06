package controller

import "github.com/labstack/echo/v4"

func History(c echo.Context) error {
	return c.File("views/history.html")
}
