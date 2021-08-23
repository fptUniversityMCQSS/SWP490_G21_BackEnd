package QA

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/entity"
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
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error020FileError,
		})
	}
	src, err := file.Open()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error021OpenFileError,
		})
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	timeNow := time.Now()
	var user = &entity.User{
		Id: intUserId,
	}
	var exam = &entity.ExamTest{
		User: user,
		Name: file.Filename,
		Date: timeNow,
	}

	i, err := utility.DB.QueryTable("exam_test").PrepareInsert()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error040CantPrepareStatementExamTest,
		})
	}
	insert, err := i.Insert(exam)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error041InsertExamFailed,
		})
	}
	err2 := i.Close()
	if err2 != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}

	userPath := "examtest/" + userName
	stringIdInsert := strconv.Itoa(int(insert))
	fileFolderPath := userPath + "/" + stringIdInsert
	err = os.MkdirAll(fileFolderPath, os.ModePerm)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error025CreateDirectoryError,
		})
	}

	filePath := fileFolderPath + "/" + file.Filename
	exam.Path = fileFolderPath
	_, err = utility.DB.Update(exam)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error042UpdateExamFailed,
		})
	}
	dst, err := os.Create(filePath)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error025CreateDirectoryError,
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
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error027CopyFileError,
		})
	}

	FileInt := c.Request().Header.Get("Content-Length")
	size, err := strconv.ParseInt(FileInt, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error043CantParseFileSize,
		})
	}

	r, err := docx.ReadDocxFromMemory(src, size)
	if err != nil {
		dir, err := filepath.Abs(fileFolderPath)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error044CantGetFilePath,
			})
		}
		fileName, err := ToDocx(dir, file.Filename)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error046CantParseFileDocToDocx,
			})
		}
		fi, err := fileName.Stat()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error047StatFileError,
			})
		}
		r, err = docx.ReadDocxFromMemory(fileName, fi.Size())
		err = fileName.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error033CloseFileError,
			})
		}
		err3 := os.Remove(fileName.Name())
		if err3 != nil {
			utility.FileLog.Println(err3)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error038RemoveFileError,
			})
		}
	}

	var xmlDocument model.XMLDocument

	err = xml.Unmarshal([]byte(r.Editable().GetContent()), &xmlDocument)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error048ParseFileXmlError,
		})
	}
	array := xmlDocument.XMLBody.XMLBodyPs

	content := ""
	for i := 0; i < len(array[0].XMLBodyPr); i++ {
		content += array[0].XMLBodyPr[i].Subject
	}
	content, err = reFormatStringToAdd(content)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error059ReformatStringFalse,
		})
	}
	exam.Subject = content
	content = ""
	for j := 0; j < len(array[1].XMLBodyPr); j++ {
		content += array[1].XMLBodyPr[j].Subject
	}
	content, err = reFormatStringToAdd(content)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error059ReformatStringFalse,
		})
	}
	numberOfQuestions, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error058ParseNumberOfQuestionsError,
		})
	}
	exam.NumberOfQuestions = numberOfQuestions

	_, err = utility.DB.Update(exam)
	tables := xmlDocument.XMLBody.XMLBodyTbls
	if tables == nil {
		err := src.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error033CloseFileError,
			})
		}
		err2 := dst.Close()
		if err2 != nil {
			utility.FileLog.Println(err2)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error033CloseFileError,
			})
		}
		dir, err := filepath.Abs("examtest/" + file.Filename)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error044CantGetFilePath,
			})
		}
		err3 := os.Remove(dir)
		if err3 != nil {
			utility.FileLog.Println(err3)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error038RemoveFileError,
			})
		}
		return c.JSON(http.StatusBadRequest, response.Message{
			Message: utility.Error050ReadFileDocOrDocxError,
		})
	}
	var Questions []*entity.Question
	keyIndex := []string{"", "A", "B", "C", "D", "E", "F"}
	for _, table := range tables {
		var QuestionModel entity.Question
		QN := ""
		Question := ""
		QuestionModel.Options = []*entity.Option{}
		for x, row := range table.XMLBodyTblR {
			var option entity.Option
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
							//if y != 0 {
							//	if content.Br.Local != "" {
							//		QuestionModel.Answer += "\n"
							//	}
							//	QuestionModel.Answer += content.QN
							//}
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
		QuestionNumber, err := strconv.ParseInt(QN, 10, 64)
		if err != nil {
			//utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error051ParseNumberOfQuestionError,
			//})
			continue
		}
		QuestionModel.Number = QuestionNumber
		QuestionModel.Content = Question
		QuestionModel.ExamTest = exam
		Questions = append(Questions, &QuestionModel)
	}
	i2, err := utility.DB.QueryTable("question").PrepareInsert()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error053CantGetQuestion,
		})
	}
	i3, err := utility.DB.QueryTable("option").PrepareInsert()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error054CantGetOption,
		})
	}
	for _, question := range Questions {
		id, err := i2.Insert(question)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error052InsertQuestionError,
			})
		}
		question.Id = id
		for _, option := range question.Options {
			_, err2 := i3.Insert(option)
			if err2 != nil {
				utility.FileLog.Println(err2)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error053InsertOptionError,
				})
			}
		}
	}
	err4 := i2.Close()
	if err4 != nil {
		utility.FileLog.Println(err4)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}
	err5 := i3.Close()
	if err5 != nil {
		utility.FileLog.Println(err5)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}
	exam.Status = "processing"
	_, err = utility.DB.Update(exam)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error042UpdateExamFailed,
		})
	}
	examResponse := response.HistoryResponse{
		Id:                insert,
		Date:              timeNow,
		Name:              exam.Name,
		User:              userName,
		Subject:           exam.Subject,
		NumberOfQuestions: exam.NumberOfQuestions,
		Status:            exam.Status,
	}
	enc := json.NewEncoder(c.Response())
	err = enc.Encode(examResponse)
	c.Response().Flush()

	res, err := utility.SendQuestions(utility.ConfigData.AIServer+"/qa", "POST", Questions)
	questionsMap := make(map[int64]*entity.Question)
	for _, question := range Questions {
		questionsMap[question.Number] = question
	}
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error055CantGetResponseModelAI})
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
			utility.FileLog.Println("Error reading HTTP response: ", err.Error())
			return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error055CantGetResponseModelAI})
		}
		str += string(b)

		if reader.Buffered() <= 0 {
			var qaResponse response.QuestionAnswerResponse
			utility.FileLog.Println(str)
			err := json.Unmarshal([]byte(str), &qaResponse)
			if err != nil {
				utility.FileLog.Println("json unmarshal from AI server failed")
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
			questionsMap[qaResponse.Qn].Answer = qaResponse.Answer
			_, err = utility.DB.Update(questionsMap[qaResponse.Qn], "answer")
			if err != nil {
				utility.FileLog.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error056UpdateAnswerError,
				})
			}
			err = enc.Encode(questionsResponse)
			if err != nil {
				utility.FileLog.Println("json encoding for responding failed")
				continue
			}
			c.Response().Flush()
			str = ""
		}
	}
	exam.Status = "finished"
	_, err = utility.DB.Update(exam)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error042UpdateExamFailed,
		})
	}
	examResponse = response.HistoryResponse{
		Id:                insert,
		Date:              timeNow,
		Name:              exam.Name,
		User:              userName,
		Subject:           exam.Subject,
		NumberOfQuestions: exam.NumberOfQuestions,
		Status:            exam.Status,
	}
	err = enc.Encode(examResponse)
	c.Response().Flush()
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
		utility.FileLog.Println(err)
		return nil, err
	}
	unknown, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	word, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	_, err = oleutil.PutProperty(word, "Visible", true)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	documents := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer documents.Release()
	document := oleutil.MustCallMethod(documents, "Open", inputDoc, false, true).ToIDispatch()
	defer document.Release()
	oleutil.MustCallMethod(document, "SaveAs", outPutDocx, 16).ToIDispatch()
	_, err = oleutil.CallMethod(word, "Quit")
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	word.Release()
	ole.CoUninitialize()
	open, err := os.Open(outPutDocx)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	return open, nil
}

func reFormatStringToAdd(s string) (string, error) {
	re, err := regexp.Compile("(.*):")
	if err != nil {
		utility.FileLog.Println(err)
		return "", err
	}
	stringAfterFormat := re.ReplaceAllString(s, "")
	return strings.TrimSpace(stringAfterFormat), nil
}
