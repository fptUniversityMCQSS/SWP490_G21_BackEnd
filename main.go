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

/*
	Require to install GCC at https://jmeubank.github.io/tdm-gcc/download/
*/
func main() {
	//start echo
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	//-----Backend-----
	backend := e.Group("/api")
	signedIn := backend.Group("", middleware.JWT([]byte(Authenticate.JwtSignature)))

	adminPermission := Role{
		"-admin-",
	}
	staffPermission := Role{
		adminPermission.requirement + "-staff-",
	}
	userPermission := Role{
		staffPermission.requirement + "-user-",
	}

	//backend.GET("/", controller.Home)
	backend.POST("/login", Authenticate.LoginResponse)
	backend.POST("/register", Authenticate.Register)

	user := signedIn.Group("", userPermission.Header)
	user.PUT("/qa", controller.QaResponse)
	user.GET("/history", controller.History)
	user.GET("/history/:id", controller.GetExamById)
	user.GET("/history/:id/download", controller.DownloadExam)
	user.GET("/knowledge", controller.ListKnowledge)

	staff := signedIn.Group("", staffPermission.Header)
	staff.PUT("/knowledge", controller.KnowledgeUpload)
	staff.GET("/knowledge/:id", controller.DownloadKnowledge)
	staff.DELETE("/knowledge/:id", controller.DeleteKnowledge)

	admin := signedIn.Group("/admin", adminPermission.Header)
	admin.GET("/user", Admin.ListUser)
	admin.POST("/user", Admin.AddUser)
	admin.GET("/user/:id", Admin.GetUserById)
	admin.DELETE("/user/:id", Admin.DeleteUserById)
	admin.PATCH("/user/:id", Admin.UpdateUser)

	//-------Frontend-------
	e.Static("/", utility.ConfigData.StaticFolder)
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.StartTLS(":"+utility.ConfigData.Port,
		utility.ConfigData.HttpsCertificate,
		utility.ConfigData.HttpsKey))
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

func customHTTPErrorHandler(err error, c echo.Context) {
	errorPage := utility.ConfigData.StaticFolder + "/index.html"
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
