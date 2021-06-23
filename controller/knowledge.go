package controller

import (
	"SWP490_G21_Backend/model/response"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"os"

	"SWP490_G21_Backend/model"

	"net/http"
)

func ListKnowledge(c echo.Context) error {

	o := orm.NewOrm()
	var knows []*model.Knowledge
	var knowRs []*response.KnowledgResponse
	// Get a QuerySeter object. User is table name
	qs, err := o.QueryTable("knowledge").RelatedSel().All(&knows)

	//if has problem in connection
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	//add selected data to knowledge_Res list
	for _, k := range knows {
		var knowR = new(response.KnowledgResponse)
		knowR.Name = k.Name
		knowR.Date = k.Date
		knowR.Username = k.User.Username
		knowRs = append(knowRs, knowR)
	}
	fmt.Printf("%d knowledges read \n", qs)
	return c.JSON(http.StatusOK, knowRs)

}

func KnowledgeUpload(c echo.Context) error {
	// Read form fields

	date := c.FormValue("date")
	userId := c.FormValue("userId")

	fmt.Println("date: ", date)
	fmt.Println("userId: ", userId)

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

	// Destination
	dst, err := os.Create("testdoc/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully ", file.Filename))

}
