package controller

import (
	"database/sql"
	"lib/model/response"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

var(
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

		_ = result.Scan(&s.Name, &s.Date , &s.Username)

		sliceUsers = append(sliceUsers, s)

	}

	return c.JSON(http.StatusOK, sliceUsers )
	//return c.Response(sliceUsers,"knowledge.html"
	//return c.File("knowledge.html")
}

func KnowledgeUpload(c echo.Context) error {
	return c.File("views/knowledgeUpload.html")
}

