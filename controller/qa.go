package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"bufio"
	"encoding/json"
	"encoding/xml"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"github.com/nguyenthenguyen/docx"
	"io"
	_ "io/ioutil"
	"log"
	"mime/multipart"
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

func QaResponse(c echo.Context) error {
	// sending file with the wrong format doesn't return error
	// not deleting if fail

	var intUserId int64
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	intUserId = int64(userId)

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "send file error",
		})
	}
	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "open file error",
		})
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	userPath := "examtest/" + userName
	timeNow := time.Now()
	re := regexp.MustCompile(":")
	timeNowAfterRegex := re.ReplaceAllString(timeNow.String(), "-")
	fileFolderPath := userPath + "/" + timeNowAfterRegex
	err = os.MkdirAll(fileFolderPath, os.ModePerm)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Can't create directory",
		})
	}

	filePath := fileFolderPath + "/" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Can't create directory of file",
		})
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Copy file error",
		})
	}

	FileInt := c.Request().Header.Get("Content-Length")
	size, err := strconv.ParseInt(FileInt, 10, 64)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Can not parse file size",
		})
	}

	r, err := docx.ReadDocxFromMemory(src, size)

	if err != nil {
		dir, err := filepath.Abs(fileFolderPath)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not get directory path",
			})
		}
		fileName, err := ToDocx(dir, file.Filename)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not parse file",
			})
		}
		fi, err := fileName.Stat()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Stat() file error",
			})
		}
		r, err = docx.ReadDocxFromMemory(fileName, fi.Size())
		err = fileName.Close()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Close file error",
			})
		}
		err3 := os.Remove(fileName.Name())
		if err3 != nil {
			log.Println(err3)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Remove file error",
			})
		}
	}

	var xmlDocument model.XMLDocument

	err = xml.Unmarshal([]byte(r.Editable().GetContent()), &xmlDocument)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Parse xml false",
		})
	}

	var user = &model.User{
		Id:       intUserId,
		Username: userName,
	}
	var exam = &model.ExamTest{
		User: user,
		Name: file.Filename,
		Date: timeNow,
		Path: filePath,
	}
	array := xmlDocument.XMLBody.XMLBodyPs
	var Content string
	for i := 0; i < 2; i++ {
		exam.Subject = Content
		NumberOfQuestions := ""
		for j := 0; j < len(array[i].XMLBodyPr); j++ {
			NumberOfQuestions += array[i].XMLBodyPr[j].Subject
		}
		re, _ := regexp.Compile("(.*)\\s")
		NumberOfQuestions = re.ReplaceAllString(NumberOfQuestions, "")
		Content = NumberOfQuestions
		exam.NumberOfQuestions, _ = strconv.ParseInt(Content, 10, 64)
	}

	i, err := utility.DB.QueryTable("exam_test").PrepareInsert()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Query table examtest failed",
		})
	}
	insert, err := i.Insert(exam)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "insert exam error",
		})
	}
	err2 := i.Close()
	if err2 != nil {
		log.Println(err2)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close connection error",
		})
	}
	tables := xmlDocument.XMLBody.XMLBodyTbls
	if tables == nil {
		err := src.Close()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Close src error",
			})
		}
		err2 := dst.Close()
		if err2 != nil {
			log.Println(err2)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Close dst error",
			})
		}
		dir, err := filepath.Abs("examtest/" + file.Filename)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not get directory path",
			})
		}
		err3 := os.Remove(dir)
		if err3 != nil {
			log.Println(err3)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Romove file error",
			})
		}
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

	i2, err := utility.DB.QueryTable("question").PrepareInsert()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Insert Question error",
		})
	}
	i3, err := utility.DB.QueryTable("option").PrepareInsert()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Insert Option error",
		})
	}
	for _, question := range Questions {
		_, err := i2.Insert(question)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Insert Option error",
			})
		}
		for _, option := range question.Options {
			_, err2 := i3.Insert(option)
			if err2 != nil {
				log.Println(err2)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: "Insert Option error",
				})
			}
		}
	}
	err4 := i2.Close()
	if err4 != nil {
		log.Println(err4)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close insert question error",
		})
	}
	err5 := i3.Close()
	if err5 != nil {
		log.Println(err5)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close insert option error",
		})
	}

	examResponse := response.HistoryResponse{
		Id:                insert,
		Date:              timeNow,
		Name:              exam.Name,
		User:              userName,
		Subject:           exam.Subject,
		NumberOfQuestions: exam.NumberOfQuestions,
	}
	enc := json.NewEncoder(c.Response())
	err = enc.Encode(examResponse)
	c.Response().Flush()

	res, err := utility.SendQuestions(utility.ConfigData.AIServer+"/qa", "POST", Questions)
	questionsMap := make(map[int64]*model.Question)
	for _, question := range Questions {
		questionsMap[question.Number] = question
	}
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{Message: "Fail to receive response from AI server"})
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	reader := bufio.NewReader(res.Body)
	str := ""
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error reading HTTP response: ", err.Error())
			return c.JSON(http.StatusInternalServerError, response.Message{Message: "Fail to receive response from AI server"})
		}
		str += string(b)

		if reader.Buffered() <= 0 {
			println(str)
			var qaResponse response.QuestionAnswerResponse
			err := json.Unmarshal([]byte(str), &qaResponse)
			if err != nil {
				log.Println("json unmarshal from AI server failed")
				continue
			}
			var optionsQAResponse []response.OptionResponse
			for _, option := range questionsMap[qaResponse.Qn].Options {
				optionsQAResponse = append(optionsQAResponse, response.OptionResponse{
					Key:     option.Key,
					Content: option.Content,
				})
			}
			var questionsResponse = response.QuestionResponse{
				Number:  qaResponse.Qn,
				Content: questionsMap[qaResponse.Qn].Content,
				Options: optionsQAResponse,
				Answer:  qaResponse.Answer,
			}
			err = enc.Encode(questionsResponse)
			if err != nil {
				log.Println("json encoding for responding failed")
				continue
			}
			c.Response().Flush()
			str = ""
		}
	}

	return c.JSON(http.StatusOK, response.Message{Message: "DONE"})

}
func RemoveEndChar(s string) string {
	sizeQuestion := len(s)

	if sizeQuestion > 0 && s[sizeQuestion-1] == '\n' {
		s = s[:sizeQuestion-1]
	}
	return s
}

func ToDocx(folderPath string, fileName string) (*os.File, error) {
	inputDoc := folderPath + "/" + fileName
	splitPath := strings.Split(fileName, ".")

	outPutDocx := ""
	for i := 0; i < len(splitPath)-1; i++ {
		outPutDocx += splitPath[i]
	}
	outPutDocx = outPutDocx + ".docx"

	outPutDocx = folderPath + "/" + outPutDocx

	err := ole.CoInitialize(0)
	if err != nil {
		log.Println(err)
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
	document := oleutil.MustCallMethod(documents, "Open", inputDoc, false, true).ToIDispatch()
	defer document.Release()
	oleutil.MustCallMethod(document, "SaveAs", outPutDocx, 16).ToIDispatch()
	_, err = oleutil.CallMethod(word, "Quit")
	if err != nil {
		return nil, err
	}
	word.Release()
	ole.CoUninitialize()
	open, err := os.Open(outPutDocx)
	if err != nil {
		return nil, err
	}
	return open, nil
}

//func convertXmlToDocx(folderPath string, fileName string) (question []*model.Question) {
//
//}
