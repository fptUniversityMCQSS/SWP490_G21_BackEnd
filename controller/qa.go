package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
)

func Qa(c echo.Context) error {
	return c.File("views/index.html")
}

func QaResponse(c echo.Context) error {

	//------------
	// Read files
	//------------

	// Multipart form
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	bs, _ := ioutil.ReadAll(src)
	fmt.Println(string(bs))
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully .</p>", file.Filename))

}
