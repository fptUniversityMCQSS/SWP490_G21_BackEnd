package main

import (
	"SWP490_G21_Backend/controller/Authenticate"
	"SWP490_G21_Backend/controller/User"
	"SWP490_G21_Backend/controller/QA"
	"SWP490_G21_Backend/controller/Knowledge"
	"SWP490_G21_Backend/controller/Log"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
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
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

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

	backend.POST("/login", (DebugHandler{Authenticate.LoginResponse}).debug)
	backend.POST("/register", (DebugHandler{Authenticate.Register}).debug)

	user := signedIn.Group("", userPermission.Header)
	user.PUT("/qa", (DebugHandler{QA.QaResponse}).debug)
	user.GET("/history", (DebugHandler{QA.History}).debug)
	user.GET("/history/:id", (DebugHandler{QA.GetExamById}).debug)
	user.DELETE("/history/:id", (DebugHandler{QA.DeleteExam}).debug)
	user.GET("/history/:id/download", (DebugHandler{QA.DownloadExam}).debug)
	user.GET("/knowledge", (DebugHandler{Knowledge.ListKnowledge}).debug)
	user.GET("/user", (DebugHandler{User.GetUserInfo}).debug)
	user.PATCH("/user", (DebugHandler{User.ChangeProfile}).debug)

	staff := signedIn.Group("", staffPermission.Header)
	staff.PUT("/knowledge", (DebugHandler{Knowledge.KnowledgeUpload}).debug)
	staff.GET("/knowledge/:id", (DebugHandler{Knowledge.DownloadKnowledge}).debug)
	staff.DELETE("/knowledge/:id", (DebugHandler{Knowledge.DeleteKnowledge}).debug)

	admin := signedIn.Group("/admin", adminPermission.Header)
	admin.GET("/user", (DebugHandler{User.ListUser}).debug)
	admin.POST("/user", (DebugHandler{User.AddUser}).debug)
	admin.GET("/user/:id", (DebugHandler{User.GetUserById}).debug)
	admin.DELETE("/user/:id", (DebugHandler{User.DeleteUserById}).debug)
	admin.PATCH("/user/:id", (DebugHandler{User.UpdateUser}).debug)
	admin.GET("/log", (DebugHandler{Log.StreamLogFile}).debug)

	//-------Frontend-------
	e.Static("/", utility.ConfigData.StaticFolder)
	e.HTTPErrorHandler = customHTTPErrorHandler

	go func() {
		e2 := echo.New()
		e2.Any("/*", redirect)
		e2.Logger.Fatal(e2.Start(":" + utility.ConfigData.PortHttp))
	}()

	e.Logger.Fatal(e.StartTLS(":"+utility.ConfigData.PortHttps,
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
		username := claims["username"].(string)
		if strings.Contains(r.requirement, role) {
			return next(c)
		} else {
			utility.FileLog.Println(username + "(role: " + role + ")" + "tried to access " + c.Path() + " . Access denied")
			return c.JSON(http.StatusForbidden, response.Message{
				Message: utility.Error057AccessDenied,
			})
		}
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	errorPage := utility.ConfigData.StaticFolder + "/index.html"
	if err := c.File(errorPage); err != nil {
		utility.FileLog.Println("static/index.html not found")
	}
}

func redirect(c echo.Context) error {
	hostParts := strings.Split(c.Request().Host, ":")
	url := "https://" + hostParts[0] + ":" + utility.ConfigData.PortHttps + c.Request().RequestURI
	return c.Redirect(http.StatusMovedPermanently, url)
}

type DebugHandler struct {
	handler echo.HandlerFunc
}

func (dHandler DebugHandler) debug(c echo.Context) error {
	DID := utility.DebugLog.Print(c.Request().Method+" "+c.Path(), false, 0)
	err := dHandler.handler(c)
	utility.DebugLog.Print(c.Request().Method+" "+c.Path(), true, DID)
	return err
}
