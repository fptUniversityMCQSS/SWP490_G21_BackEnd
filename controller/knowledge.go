package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"lib/model"
	"net/http"

)

//func init(){
//	orm.RegisterDriver("mysql",orm.DRMySQL)
//	err := orm.RegisterDataBase("default","mysql","root:abcd@/testdb?charset=utf8")
//	if err != nil{
//		glog.Fatal("Failed to register database %v",err)
//	}
//}
func Knowledge(c echo.Context) error {
	db, err := sql.Open("mysql", "root:abcd@tcp(127.0.0.1:3306)/testdb")
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


	var sliceUsers []model.Knowledge
	result, _ := db.Query("SELECT id , name , date FROM testdb.knowledge")
	for result.Next() {
		var s model.Knowledge
		_ = result.Scan(&s.ID,&s.Name, &s.Date)
		sliceUsers = append(sliceUsers, s)
	}

	return c.JSON(http.StatusOK, sliceUsers)
	//return c.File("views/knowledge.html",response)
}

func KnowledgeUpload(c echo.Context) error {
	return c.File("views/knowledgeUpload.html")
}

