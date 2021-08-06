package main

import (
	"SWP490_G21_Backend/controller"
	"SWP490_G21_Backend/controller/Admin"
	"SWP490_G21_Backend/controller/Authenticate"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func main() {
	//start echo
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	signedIn := e.Group("", middleware.JWT([]byte(Authenticate.JwtSignature)))

	adminPermission := Role{
		"-admin-",
	}
	staffPermission := Role{
		adminPermission.requirement + "-staff-",
	}
	userPermission := Role{
		staffPermission.requirement + "-user-",
	}

	//api
	e.GET("/", controller.Home)
	e.GET("/api", controller.ApiWeb)
	e.POST("/login", Authenticate.LoginResponse)
	e.POST("/register", Authenticate.Register)

	user := signedIn.Group("", userPermission.Header)
	user.PUT("/qa", controller.QaResponse)
	user.GET("/history", controller.History)
	user.GET("/history/:id", controller.GetExamById)
	user.GET("/history/:id/download", controller.DownloadExam)

	staff := signedIn.Group("", staffPermission.Header)
	staff.GET("/knowledge", controller.ListKnowledge)
	staff.PUT("/knowledge", controller.KnowledgeUpload)
	staff.GET("/knowledge/:id", controller.DownloadKnowledge)
	staff.DELETE("/knowledge/:id", controller.DeleteKnowledge)

	admin := signedIn.Group("/admin", adminPermission.Header)
	admin.GET("/user", Admin.ListUser)
	admin.POST("/user", Admin.AdminAddUser)
	admin.GET("/user/:id", Admin.GetUserById)
	admin.DELETE("/user/:id", Admin.DeleteUserById)
	admin.PATCH("/user/:id", Admin.UpdateUser)

	e.Logger.Fatal(e.Start(":" + utility.ConfigData.PortBackend))

}

type Role struct {
	requirement string
}

func (r *Role) Header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if strings.Contains(r.requirement, role) {
			return next(c)
		} else {
			return c.JSON(http.StatusForbidden, response.Message{
				Message: "Access denied",
			})
		}
	}
}
