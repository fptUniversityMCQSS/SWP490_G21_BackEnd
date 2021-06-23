package controller

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Home(c echo.Context) error {
	sess, _ := session.Get("session", c)
	user, ok := sess.Values["username"].(string)
	if ok {
		log.Println("true")
	} else {
		log.Println("false")
	}
	return c.HTML(http.StatusOK, "User: "+user)
}
