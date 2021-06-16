package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func Upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	data, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	fmt.Println("Contents of file:", string(data))

	//// Destination
	//dst, err := os.Create(file.Filename)
	//if err != nil {
	//	return err
	//}
	//defer dst.Close()

	//// Copy
	//if _, err = io.Copy(dst, src); err != nil {
	//	return err
	//}

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}

