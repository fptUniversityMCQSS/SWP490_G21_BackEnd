package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
	"os"
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
	bs, _ := ioutil.ReadAll(src)
	fmt.Println(string(bs))
	defer src.Close()

	dst, err := os.Create("testdoc/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully .</p>", file.Filename))

}
