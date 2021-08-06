package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func ListKnowledge(c echo.Context) error {

	o := orm.NewOrm()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

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
	log.Printf(userName + " get list knowledge")
	return c.JSON(http.StatusOK, knowRs)

}

func KnowledgeUpload(c echo.Context) error {
	// Read form fields
	o := orm.NewOrm()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
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

	userPath := "knowledge/" + userName
	timeNow := time.Now()
	re := regexp.MustCompile(":")
	timeNowAfterRegex := re.ReplaceAllString(timeNow.String(), "-")
	fileFolderPath := userPath + "/" + timeNowAfterRegex
	if _, err := os.Stat(userPath); os.IsNotExist(err) {
		err := os.Mkdir(userPath, os.ModeDir)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	if _, err := os.Stat(fileFolderPath); os.IsNotExist(err) {
		err := os.Mkdir(fileFolderPath, os.ModeDir)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	filePath := fileFolderPath + "/" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		return err
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
		Username: userName,
	}

	know := &model.Knowledge{
		Name: file.Filename,
		User: user,
		Path: fileFolderPath,
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

	//utility.SendFileRequest(utility.ConfigData.AIServer+"/knowledge", "POST", "knowledge/"+file.Filename)

	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   "Ready",
	}

	log.Printf(userName + " upload file : " + file.Filename)
	return c.JSON(http.StatusOK, knowledgeResponse)
}

func DownloadKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

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

	//err2, _ := os.Create("knowledgeknowledge/" + knowledge.Name + ".txt")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//return c.JSON(http.StatusOK, knowledge)
	log.Printf(userName + " downloaded " + knowledge.Name)

	return c.Attachment(knowledge.Path+"/"+knowledge.Name, knowledge.Name)
}

func DeleteKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

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

	err2 := os.RemoveAll(knowledge.Path)

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "remove file error",
		})
	}
	message := response.Message{
		Message: "Delete successfully",
	}
	log.Printf(userName + " delete file " + knowledge.Name)
	return c.JSON(http.StatusOK, message)
}
