package controller

import (
	"SWP490_G21_Backend/model/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, response.Message{Message: "Server is running..."})
}
