package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"github.com/nguyenthenguyen/docx"
	"io"
	_ "io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	_ "unsafe"
)

func Qa(c echo.Context) error {
	return c.File("views/index.html")
}

func QaResponse(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {

		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("examtest/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//examtest,err := ConvertXmlToExamTest(src)
	//if err != nil {
	//	  convertedDocx ,err := ToDocx()
	//	  if err != nil {
	//	  	log.Println(err)
	//	  }
	//	examtest,err = ConvertXmlToExamTest(convertedDocx)
	//	if err != nil {
	//		log.Println(err)
	//		//xac dinh file vo van
	//	}
	//}

	//
	//if err != nil {
	//	// Could not obtain stat, handle error
	//}
	//
	//fmt.Printf("The file is %d bytes long", fi.Size())
	FileInt := c.Request().Header.Get("Content-Length")
	size, err := strconv.ParseInt(FileInt, 10, 64)
	if err != nil {
		log.Println(err)
	}
	r, err := docx.ReadDocxFromMemory(src, size)
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(dat
	//a io.ReaderAt, size int64)
	if err != nil {
		dir, err := filepath.Abs("examtest/" + file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		fileName, err := ToDocx(dir)
		if err != nil {
			log.Fatal(err)
		}
		fi, err := fileName.Stat()
		if err != nil {
			log.Fatal(err)
		}
		r, err = docx.ReadDocxFromMemory(fileName, fi.Size())
		fileName.Close()
		os.Remove(fileName.Name())

		if err != nil {
			log.Fatal(err)
		}
	}

	var xmlDocument model.XMLDocument

	err = xml.Unmarshal([]byte(r.Editable().GetContent()), &xmlDocument)
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
		Name: file.Filename,
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
	if err != nil {
		log.Println(err)
	}
	insert, err := i.Insert(exam)
	if err != nil {
		log.Println(err)
	}
	i.Close()
	tables := xmlDocument.XMLBody.XMLBodyTbls
	if tables == nil {
		src.Close()
		dst.Close()
		dir, err := filepath.Abs("examtest/" + file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		os.Remove(dir)
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: "Cant not read file doc or docx please try again",
		})
	}
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
		QuestionModel.ExamTest = exam
		Questions = append(Questions, &QuestionModel)
	}

	i2, err := o.QueryTable("question").PrepareInsert()
	i3, err := o.QueryTable("option").PrepareInsert()
	for _, question := range Questions {
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

func ToDocx(path string) (*os.File, error) {
	filename2 := path[:len(path)-4] + ".docx"
	err := ole.CoInitialize(0)
	if err != nil {
		return nil, err
	}
	unknown, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		return nil, err
	}
	word, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	_, err = oleutil.PutProperty(word, "Visible", true)
	if err != nil {
		return nil, err
	}
	documents := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer documents.Release()
	document := oleutil.MustCallMethod(documents, "Open", path, false, true).ToIDispatch()
	defer document.Release()
	oleutil.MustCallMethod(document, "SaveAs", filename2, 16).ToIDispatch()
	//_, err = oleutil.PutProperty(document, "Saved", true)
	//if err != nil {
	//	return nil,err
	//}
	//_, err = oleutil.CallMethod(documents, "Close", false)
	//if err != nil {
	//	return nil,err
	//}
	_, err = oleutil.CallMethod(word, "Quit")
	if err != nil {
		return nil, err
	}
	word.Release()
	ole.CoUninitialize()
	open, err := os.Open(filename2)
	if err != nil {
		return nil, err
	}
	return open, nil
}
