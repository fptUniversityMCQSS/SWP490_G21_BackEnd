package controller

import (
	"SWP490_G21_Backend/model/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ApiWeb(c echo.Context) error {
	return c.JSON(http.StatusOK, response.Message{Message: "Still on development"})
}
