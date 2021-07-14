package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"github.com/nguyenthenguyen/docx"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Qa(c echo.Context) error {
	return c.File("views/index.html")
}

func QaResponse(c echo.Context) error {

	//var questions = []*model.Question{
	//	{
	//		Content:  "Kha Banh di tu bao nhieu nam",
	//		Options:  nil,
	//		Answer:   "A",
	//		ExamTest: exam,
	//		Number:   1,
	//		Mark:     1,
	//	},
	//	{
	//		Content:  "Kha Banh bao nhieu tuoi",
	//		Options:  nil,
	//		Answer:   "B",
	//		ExamTest: exam,
	//		Number:   2,
	//		Mark:     1,
	//	},
	//}
	//var option1 = []*model.Option{
	//	{
	//		Id:         1,
	//		QuestionId: questions[0],
	//		Key:        "A",
	//		Content:    "10 nam",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         2,
	//		QuestionId: questions[0],
	//		Key:        "B",
	//		Content:    "9 nam",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         3,
	//		QuestionId: questions[0],
	//		Key:        "C",
	//		Content:    "8 nam",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         4,
	//		QuestionId: questions[0],
	//		Key:        "D",
	//		Content:    "khong bi di tu",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//}
	//var option2 = []*model.Option{
	//	{
	//		Id:         5,
	//		QuestionId: questions[1],
	//		Key:        "A",
	//		Content:    "20 t",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         6,
	//		QuestionId: questions[1],
	//		Key:        "B",
	//		Content:    "21 t",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         7,
	//		QuestionId: questions[1],
	//		Key:        "C",
	//		Content:    "22 t",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//	{
	//		Id:         8,
	//		QuestionId: questions[1],
	//		Key:        "D",
	//		Content:    "23 t",
	//		Paragraph:  "Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam",
	//	},
	//}
	//
	//var options = [][]*model.Option{
	//	option1, option2,
	//}
	//
	//for i, question := range questions {
	//	question.Options = options[i]
	//}

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
	r, err := docx.ReadDocxFile("testdoc/" + file.Filename)
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create("parseDocxToXML/" + file.Filename + ".xml")
	defer f.Close()
	f.WriteString(r.Editable().GetContent())

	xmlFile, err := os.Open("parseDocxToXML/" + file.Filename + ".xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var xmlDocument model.XMLDocument
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &xmlDocument)
	if err != nil {
		fmt.Println(err)
	}

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
	array := xmlDocument.XMLBody.XMLBodyPs
	var VatChua2 string
	for i := 0; i < 2; i++ {
		exam.Subject = VatChua2
		VatChua := ""
		for j := 0; j < len(array[i].XMLBodyPr); j++ {
			VatChua += array[i].XMLBodyPr[j].Subject
		}
		re, _ := regexp.Compile("(.*)\\s")
		VatChua = re.ReplaceAllString(VatChua, "")
		VatChua2 = VatChua
		exam.NumberOfQuestions, _ = strconv.ParseInt(VatChua2, 10, 64)
	}
	i, err := o.QueryTable("exam_test").PrepareInsert()
	insert, err := i.Insert(exam)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(insert)
	i.Close()
	tables := xmlDocument.XMLBody.XMLBodyTbls
	var Questions []*model.Question
	keyIndex := []string{"", "A", "B", "C", "D", "E", "F"}
	for _, table := range tables {
		var QuestionModel model.Question
		QN := ""
		Question := ""
		QuestionModel.Options = []*model.Option{}
		for x, row := range table.XMLBodyTblR {
			var option model.Option
			for y, column := range row.XMLBodyTblRC {
				for _, paragraph := range column.XMLBodyTblRCP {
					for _, content := range paragraph.XMLBodyTblRCPR {
						if x == 0 {
							if y == 0 {
								QN += content.QN
							} else {
								if content.Br.Local != "" {
									Question += "\n"
								}
								Question += content.QN
							}
						} else if x <= 6 {
							if y != 0 {
								if content.Br.Local != "" {
									option.Content += "\n"
								}
								option.Content += content.QN
							}
						} else if x == 7 {
							if y != 0 {
								QuestionModel.Answer = content.QN
							}
						}
					}
					if y != 0 {
						if x == 0 {
							Question += "\n"
						} else if x <= 6 {
							option.Content += "\n"
						}
					}
				}
			}
			if x > 0 && x <= 6 {
				option.Content = RemoveEndChar(option.Content)
				option.Content = strings.TrimSpace(option.Content)
				if option.Content != "" {
					option.Key = keyIndex[x]
					option.QuestionId = &QuestionModel
					QuestionModel.Options = append(QuestionModel.Options, &option)
				}
			}
		}
		Question = RemoveEndChar(Question)
		re, _ := regexp.Compile("(.*)=")
		QN = re.ReplaceAllString(QN, "")
		QuestionNumber, _ := strconv.ParseInt(QN, 10, 64)
		QuestionModel.Number = QuestionNumber
		QuestionModel.Content = Question
		Questions = append(Questions, &QuestionModel)
	}
	i2, err := o.QueryTable("question").PrepareInsert()
	i3, err := o.QueryTable("option").PrepareInsert()
	for _, question := range Questions {
		question.ExamTest = exam
		id, err := i2.Insert(question)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id)
		for _, option := range question.Options {
			id2, err2 := i3.Insert(option)
			if err2 != nil {
				log.Println(err2)
			}
			fmt.Println(id2)
		}
	}
	i2.Close()
	i3.Close()

	examResponse := response.HistoryResponse{
		Id:   insert,
		Name: exam.Name,
		Date: time.Now(),
	}

	return c.JSON(http.StatusOK, examResponse)

}
func RemoveEndChar(s string) string {
	sizeQuestion := len(s)

	if sizeQuestion > 0 && s[sizeQuestion-1] == '\n' {
		s = s[:sizeQuestion-1]
	}
	return s
}
