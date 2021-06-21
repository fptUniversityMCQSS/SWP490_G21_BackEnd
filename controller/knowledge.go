package controller

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"lib/model"
	"lib/model/response"
	"net/http"
	"os"
)

var (
	db, err = sql.Open("mysql", "root:abcd@tcp(127.0.0.1:3306)/testdb")
)

func Knowledge(c echo.Context) error {

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	var sliceUsers []response.KnowledgeRes
	result, _ := db.Query("select k.name , k.date , u.username \nfrom knowledge as k , user as u \nwhere k.user_id = u.id")
	for result.Next() {
		var s response.KnowledgeRes

		_ = result.Scan(&s.Name, &s.Date, &s.Username)

		sliceUsers = append(sliceUsers, s)

	}

	return c.JSON(http.StatusOK, sliceUsers)
	//return c.Response(sliceUsers,"knowledge.html"
	//return c.File("knowledge.html")
}

func KnowledgeUpload(c echo.Context) error {
	// Read form fields

	date := c.FormValue("date")
	userId := c.FormValue("userId")

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

	// add knowledge to database
	k := &model.Knowledge{}
	if err := c.Bind(k); err != nil {
		return err
	}
	insertDB, err := db.Prepare("INSERT INTO knowledge(name, date,userId) values (?,?,?);")
	if err != nil {
		panic(err.Error())
	}
	exec, err := insertDB.Exec(file.Filename, date, userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exec)

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully ", file.Filename))
}
