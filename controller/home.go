package controller

import (
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.File("views/home.html")
}
