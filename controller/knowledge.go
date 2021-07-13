package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"fmt"
	"strconv"
)

func ListKnowledge(c echo.Context) error {

	o := orm.NewOrm()
	var knows []*model.Knowledge
	var knowRs []*response.KnowledgResponse
	// Get a QuerySeter object. User is table name
	_, err := o.QueryTable("knowledge").OrderBy("-id").RelatedSel().All(&knows)

	//if has problem in connection
	if err != nil {
		return err
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
	//fmt.Println("date: ", date)
	//fmt.Println("userId: ", userId)

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

	//defer func(src multipart.File) {
	//	err := src.Close()
	//	if err != nil {
	//
	//	}
	//}(src)
	//data, err := ioutil.ReadAll(src)
	//if err != nil {
	//	fmt.Println("File reading error", err)
	//	return err
	//}
	//fmt.Println("Contents of file:", string(data))

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

	//qs := o.QueryTable("knowledge")

	//dateParsed, err := time.Parse("Jan 2, 2006 at 3:04pm (MST)", time.Now().String())
	//if err != nil {
	//	fmt.Println(err)
	//}
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
		return err
	}
	//fmt.Println(i)
	insert, err := i.Insert(know)
	if err != nil {
		return err
	}
	//fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return err1
	}

	knowledgeResponse := response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
	}

	return c.JSON(http.StatusOK, knowledgeResponse)
}

func DownloadKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	knowledgeId := c.Param("id")
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	var knowledge model.Knowledge

	err := o.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)

	//if has problem in connection
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}

	//err2, _ := os.Create("testdoc/" + knowledge.Name + ".txt")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//return c.JSON(http.StatusOK, knowledge)

	return c.Attachment("testdoc/"+knowledge.Name, knowledge.Name)
}

func DeleteKnowledge(c echo.Context) error {
	o := orm.NewOrm()
	knowledgeId := c.Param("id")
	var knowledge model.Knowledge
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	err3 := o.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)
	if err3 != nil {
		fmt.Println(err3)
	}
	_, err := o.QueryTable("knowledge").Filter("id", intKnowledgeId).Delete()
	if err != nil {
		fmt.Println(err)
	}
	err2 := os.Remove("testdoc/" + knowledge.Name)

	if err2 != nil {
		fmt.Println(err2)

	}
	return c.JSON(http.StatusOK, "Delete successfully")
}
