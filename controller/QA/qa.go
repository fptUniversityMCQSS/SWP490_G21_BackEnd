package QA

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/entity"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"bufio"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/gommon/log"
	"github.com/nguyenthenguyen/docx"
	"io"
	"io/ioutil"
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

var RequestingQA = make(map[int64]chan error)

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
	resultError := ""
	exam.Questions, resultError = processQuestion(exam, size, dst, src, fileFolderPath, file)
	if resultError != "" {
		_, err = utility.DB.Delete(exam)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error062DeleteExamFailed,
			//})
		}
		utility.FileLog.Println(resultError)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: resultError,
		})
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
	for _, question := range exam.Questions {
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

	done := make(chan error)
	RequestingQA[exam.Id] = done
	cx, cancel := context.WithCancel(context.Background())
	messageChan := make(chan string)
	go func() {
		res, err := utility.SendQuestions(utility.ConfigData.AIServer+"/qa", "POST", exam.Questions, &cx)
		if err != nil {
			cancel()
			delete(RequestingQA, exam.Id)
			utility.FileLog.Println(err)
			done <- response.Message{Message: utility.Error055CantGetResponseModelAI}
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
					done <- nil
					return
				}
				utility.FileLog.Println("Error reading HTTP response: ", err.Error())
				done <- response.Message{Message: utility.Error055CantGetResponseModelAI}
				return
			}
			str += string(b)

			if reader.Buffered() <= 0 {
				messageChan <- str
				str = ""
			}
		}
	}()
	questionsMap := make(map[int64]*entity.Question)
	for _, question := range exam.Questions {
		questionsMap[question.Number] = question
	}
	err = nil
	for {
		flag := false
		select {
		case <-cx.Done():
			flag = true
			//break
		case err = <-done:
			flag = true
			//break
		case str := <-messageChan:
			err := func() error {
				var qaResponse response.QuestionAnswerResponse
				err := json.Unmarshal([]byte(str), &qaResponse)
				if err != nil {
					utility.FileLog.Println("json unmarshal from AI server failed: " + err.Error())
					return nil
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
					return err
				}
				err = enc.Encode(questionsResponse)
				if err != nil {
					utility.FileLog.Println("json encoding for responding failed")
					return nil
				}
				c.Response().Flush()
				return nil
			}()
			if err != nil {
				cancel()
				delete(RequestingQA, exam.Id)
				utility.FileLog.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error056UpdateAnswerError,
				})
			}
			//break
		}
		if flag {
			utility.FileLog.Println(err)
			break
		}
	}
	cancel()
	delete(RequestingQA, exam.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{Message: err.Error()})
	}
	exam.Status = "finished"
	_, err = utility.DB.Update(exam)
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
func processQuestion(exam *entity.ExamTest, size int64, dst *os.File, src multipart.File, fileFolderPath string, file *multipart.FileHeader) ([]*entity.Question, string) {

	r, err := docx.ReadDocxFromMemory(src, size)
	if err != nil {
		dir, err := filepath.Abs(fileFolderPath)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error044CantGetFilePath,
			//})
			return nil, utility.Error044CantGetFilePath
		}
		fileName, err := ToDocx(dir, file.Filename)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error046CantParseFileDocToDocx,
			//})
			return nil, utility.Error046CantParseFileDocToDocx
		}
		fi, err := fileName.Stat()
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error047StatFileError,
			//})
			return nil, utility.Error047StatFileError
		}
		r, err = docx.ReadDocxFromMemory(fileName, fi.Size())
		err = fileName.Close()
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error033CloseFileError,
			//})
			return nil, utility.Error033CloseFileError
		}
		err3 := os.Remove(fileName.Name())
		if err3 != nil {
			utility.FileLog.Println(err3)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error038RemoveFileError,
			//})
			return nil, utility.Error038RemoveFileError
		}
	}

	var xmlDocument model.XMLDocument

	err = xml.Unmarshal([]byte(r.Editable().GetContent()), &xmlDocument)
	if err != nil {
		utility.FileLog.Println(err)
		//return c.JSON(http.StatusInternalServerError, response.Message{
		//	Message: utility.Error048ParseFileXmlError,
		//})
		return nil, utility.Error048ParseFileXmlError
	}
	array := xmlDocument.XMLBody.XMLBodyPs
	if len(array) >= 2 {
		fmt.Println(len(array))
		content := ""
		for i := 0; i < len(array[0].XMLBodyPr); i++ {
			content += array[0].XMLBodyPr[i].Subject
		}
		content, err = reFormatStringToAdd(content)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error059ReformatStringFalse,
			//})
			//return nil, utility.Error059ReformatStringFalse
		}
		exam.Subject = content
		content = ""
		for j := 0; j < len(array[1].XMLBodyPr); j++ {

			content += array[1].XMLBodyPr[j].Subject
		}
		content, err = reFormatStringToAdd(content)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error059ReformatStringFalse,
			//})
			//return nil, utility.Error059ReformatStringFalse
		}
		numberOfQuestions, err := strconv.ParseInt(content, 10, 64)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error058ParseNumberOfQuestionsError,
			//})
			//return nil, utility.Error058ParseNumberOfQuestionsError
		}
		exam.NumberOfQuestions = numberOfQuestions
	}
	_, err = utility.DB.Update(exam)
	tables := xmlDocument.XMLBody.XMLBodyTbls
	if tables == nil {
		err := src.Close()
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error033CloseFileError,
			//})
			return nil, utility.Error033CloseFileError
		}
		err2 := dst.Close()
		if err2 != nil {
			utility.FileLog.Println(err2)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error033CloseFileError,
			//})
			return nil, utility.Error033CloseFileError
		}
		dir, err := filepath.Abs("examtest/" + file.Filename)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error044CantGetFilePath,
			//})
			return nil, utility.Error044CantGetFilePath
		}
		err3 := os.Remove(dir)
		if err3 != nil {
			utility.FileLog.Println(err3)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error038RemoveFileError,
			//})
			return nil, utility.Error038RemoveFileError
		}
		_, err = utility.DB.Delete(exam)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: utility.Error062DeleteExamFailed,
			//})
			return nil, utility.Error062DeleteExamFailed
		}
		//return c.JSON(http.StatusBadRequest, response.Message{
		//	Message: utility.Error050ReadFileDocOrDocxError,
		//})
		return nil, utility.Error050ReadFileDocOrDocxError
	}
	var Questions []*entity.Question

	for _, table := range tables {
		var isOptions []bool
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
						} else {
							if y == 0 {
								matchReg, _ := regexp.MatchString("^\\W*\\w{1}\\W*$", content.QN)
								isOptions = append(isOptions, matchReg)
								if matchReg == true {
									reString := regexp.MustCompile("\\W")
									content.QN = reString.ReplaceAllString(content.QN, "")
									option.Key = content.QN
								}
							}
							if y != 0 && isOptions != nil && isOptions[x-1] {
								if content.Br.Local != "" {
									option.Content += "\n"
								}
								option.Content += content.QN
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
			if x > 0 && isOptions != nil && isOptions[x-1] {
				option.Content = RemoveEndChar(option.Content)
				option.Content = strings.TrimSpace(option.Content)
				if option.Content != "" {
					option.QuestionId = &QuestionModel
					QuestionModel.Options = append(QuestionModel.Options, &option)
				}
			}
		}
		Question = RemoveEndChar(Question)
		QNOfQuestions := ""
		re, _ := regexp.Compile("\\d+")
		arrayOfQN := re.FindAllString(QN, -1)
		if len(arrayOfQN) == 0 {

		} else {
			QNOfQuestions = arrayOfQN[0]
		}
		QuestionNumber, err := strconv.ParseInt(QNOfQuestions, 10, 64)
		if err != nil {
			utility.FileLog.Println(err)
			//return c.JSON(http.StatusInternalServerError, response.Message{
			//	Message: ,
			//})
			continue
		}
		QuestionModel.Number = QuestionNumber
		QuestionModel.Content = Question
		QuestionModel.ExamTest = exam
		if strings.TrimSpace(QuestionModel.Content) == "" {
			continue
		}
		if QuestionModel.Options == nil || len(QuestionModel.Options) == 0 {
			continue
		}
		Questions = append(Questions, &QuestionModel)
	}

	return Questions, ""
}

func QaGenerateDocx(c echo.Context) error {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	var examTest entity.ExamTest
	err := json.NewDecoder(c.Request().Body).Decode(&examTest)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error070ParseJsonError,
		})
	}
	fmt.Println(examTest.Questions[0].Content)
	utility.FileLog.Println(userName + " create " + examTest.Name)
	formatFile, err := formatFileDocxFromRequestBody(examTest)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error045CantGetResult,
		})
	}
	return c.Attachment(formatFile, examTest.Name)
}

func formatFileDocxFromRequestBody(exam entity.ExamTest) (string, error) {
	r, err := docx.ReadDocxFile("template/TestFormat.docx")

	table, err := ioutil.ReadFile("template/table.xml")
	if err != nil {
		return err.Error(), err
	}
	optionTable, err := ioutil.ReadFile("template/option.xml")
	if err != nil {
		return err.Error(), err
	}
	body, err := ioutil.ReadFile("template/body.xml")
	if err != nil {
		return err.Error(), err
	}
	tableContent := ""

	for num, question := range exam.Questions {
		newTable := strings.ReplaceAll(string(table), "{{numberOfQuestion}}", writeStringDocx(strconv.Itoa(num+1)))
		newTable = strings.ReplaceAll(newTable, "{{QuestionContent}}", writeStringDocx(question.Content))
		optionContent := ""
		for _, option := range question.Options {
			newOption := strings.ReplaceAll(string(optionTable), "{{optionKey}}", writeStringDocx(option.Key))
			newOption = strings.ReplaceAll(newOption, "{{optionContent}}", writeStringDocx(option.Content))
			optionContent += newOption
		}
		newTable = strings.ReplaceAll(newTable, "{{option}}", optionContent)
		newTable = strings.ReplaceAll(newTable, "{{answers}}", writeStringDocx(question.Answer))
		tableContent += newTable
	}
	bodyContent := strings.ReplaceAll(string(body), "{{subject}}", writeStringDocx(exam.Subject))
	bodyContent = strings.ReplaceAll(bodyContent, "{{numberOfQuestions}}", writeStringDocx(strconv.Itoa(int(exam.NumberOfQuestions))))
	bodyContent = strings.ReplaceAll(bodyContent, "{{table}}", tableContent)
	editFile := r.Editable()
	editFile.SetContent(bodyContent)
	pathToSave := "template/Temp.docx"
	err = editFile.WriteToFile(pathToSave)
	if err != nil {
		return err.Error(), err
	}
	return pathToSave, nil
}
