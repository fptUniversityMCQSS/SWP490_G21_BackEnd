package controller

import (
	"SWP490_G21_Backend/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"os"
	"strings"
)

func Qa(c echo.Context) error {
	return c.File("views/index.html")
}

func QaResponse(c echo.Context) error {
	o := orm.NewOrm()
	var intUserId int64
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64) //Ep kieu sang float64
	intUserId = int64(userId)
	var user = &model.User{
		Id: intUserId,
	}
	var exam = &model.ExamTest{
		User: user,
		Name: "oke",
	}
	var questions = []*model.Question{
		{
			Content:  "Kha Banh di tu bao nhieu nam",
			Options:  nil,
			Answer:   "A",
			ExamTest: exam,
			Number:   1,
			Mark:     1,
		},
		{
			Content:  "Kha Banh bao nhieu tuoi",
			Options:  nil,
			Answer:   "B",
			ExamTest: exam,
			Number:   2,
			Mark:     1,
		},
	}
	var option1 = []*model.Option{
		{
			Id:         1,
			QuestionId: questions[0],
			Key:        "A",
			Content:    "10 nam",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         2,
			QuestionId: questions[0],
			Key:        "B",
			Content:    "9 nam",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         3,
			QuestionId: questions[0],
			Key:        "C",
			Content:    "8 nam",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         4,
			QuestionId: questions[0],
			Key:        "D",
			Content:    "khong bi di tu",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
	}
	var option2 = []*model.Option{
		{
			Id:         5,
			QuestionId: questions[1],
			Key:        "A",
			Content:    "20 t",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         6,
			QuestionId: questions[1],
			Key:        "B",
			Content:    "21 t",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         7,
			QuestionId: questions[1],
			Key:        "C",
			Content:    "22 t",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
		{
			Id:         8,
			QuestionId: questions[1],
			Key:        "D",
			Content:    "23 t",
			Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
		},
	}

	var options = [][]*model.Option{
		option1, option2,
	}

	for i, question := range questions {
		question.Options = options[i]
	}

	// Parse to int64

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

	i, err := o.QueryTable("exam_test").PrepareInsert()
	i2, err := o.QueryTable("question").PrepareInsert()
	i3, err := o.QueryTable("option").PrepareInsert()
	if err != nil {
		return err
	}

	fmt.Println(i)
	insert, err := i.Insert(exam)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(insert)
	i.Close()

	for _, question := range questions {
		id, err := i2.Insert(question)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id)

	}

	for _, option := range options {
		for _, o := range option {
			id2, err2 := i3.Insert(o)
			if err2 != nil {
				log.Println(err2)
			}
			fmt.Println(id2)
		}
	}

	i2.Close()
	i3.Close()
	//err1 := i.Close()
	//if err1 != nil {
	//	return err1
	//}
	//insert2, err2 := i2.Insert(questions)
	//if err2 != nil {
	//	log.Println(err2)
	//}
	////fmt.Println(insert)
	//fmt.Println(insert2)
	//err3 := i2.Close()
	//if err3 != nil {
	//	return err3
	//}
	//insert3, err4 := i3.Insert(option1)
	//if err4 != nil {
	//	log.Println(err4)
	//}
	////fmt.Println(insert)
	//fmt.Println(insert3)

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully .</p>", file.Filename))

}
