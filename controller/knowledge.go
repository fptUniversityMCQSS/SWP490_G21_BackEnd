package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ListKnowledge(c echo.Context) error {

	o := orm.NewOrm()
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	//userId := claims["userId"].(float64)

	var knows []*model.Knowledge
	var knowRs []*response.KnowledgResponse
	// Get a QuerySeter object. User is table name
	_, err := o.QueryTable("knowledge").OrderBy("-id").RelatedSel().All(&knows)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}
	//add selected data to knowledge_Res list
	for _, k := range knows {
		var knowR = new(response.KnowledgResponse)
		knowR.Id = k.Id
		knowR.Name = k.Name
		knowR.Date = k.Date
		knowR.Username = k.User.Username
		knowRs = append(knowRs, knowR)
	}
	log.Printf(username + "get list knowledge")
	return c.JSON(http.StatusOK, knowRs)

}

func KnowledgeUpload(c echo.Context) error {
	// Read form fields
	o := orm.NewOrm()
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	username := claims["username"].(string)
	IntUserId := int64(userId)

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "file error",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "open error",
		})
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("testdoc/" + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Create file in testdoc/ error",
		})
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Copy file error",
		})
	}

	user := &model.User{
		Id:       IntUserId,
		Username: username,
	}

	know := &model.Knowledge{
		Name: file.Filename,
		User: user,
	}
	i, err := o.QueryTable("knowledge").PrepareInsert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Query error",
		})
	}
	//fmt.Println(i)
	insert, err := i.Insert(know)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Insert file error",
		})
	}
	//fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close error",
		})
	}

	knowledgeResponse := response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   "Encoding",
	}

	enc := json.NewEncoder(c.Response())
	enc.Encode(knowledgeResponse)
	c.Response().Flush()

	time.Sleep(3 * time.Second)

	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   "Ready",
	}

	log.Printf(username + " upload file : " + file.Filename)
	return c.JSON(http.StatusOK, knowledgeResponse)
}

func DownloadKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	knowledgeId := c.Param("id")
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	var knowledge model.Knowledge

	err := o.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}

	//err2, _ := os.Create("testdoc/" + knowledge.Name + ".txt")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//return c.JSON(http.StatusOK, knowledge)
	log.Printf(username + " downloaded " + knowledge.Name)
	return c.Attachment("testdoc/"+knowledge.Name, knowledge.Name)
}

func DeleteKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	knowledgeId := c.Param("id")
	var knowledge model.Knowledge
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	err3 := o.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}
	_, err := o.QueryTable("knowledge").Filter("id", intKnowledgeId).Delete()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}
	err2 := os.Remove("testdoc/" + knowledge.Name)

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "remove file error",
		})
	}
	message := response.Message{
		Message: "Delete successfully",
	}
	log.Printf(username + " delete file " + knowledge.Name)
	return c.JSON(http.StatusOK, message)
}
