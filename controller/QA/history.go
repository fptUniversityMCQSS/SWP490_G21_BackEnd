package QA

import (
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/model/unity"
	"SWP490_G21_Backend/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nguyenthenguyen/docx"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func History(c echo.Context) error {
	var history []*unity.ExamTest
	var hist []*response.HistoryResponse

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)

	_, err := utility.DB.QueryTable("exam_test").Filter("user_id", userId).All(&history)

	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error014ErrorQueryForGetAllExamTest,
		})
	}

	//add selected data to knowledge_Res list
	for _, h := range history {
		var his = &response.HistoryResponse{
			Id:                h.Id,
			Date:              h.Date,
			Name:              h.Name,
			User:              h.User.Username,
			Subject:           h.Subject,
			NumberOfQuestions: h.NumberOfQuestions,
		}

		hist = append(hist, his)
	}
	utility.FileLog.Println(userName + " get list history")
	if hist == nil {
		hist = []*response.HistoryResponse{}
	}
	return c.JSON(http.StatusOK, hist)

}
func GetExamById(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error008UserIdInvalid,
		})
	}
	var examTest unity.ExamTest
	var user unity.User
	var questionAll []*unity.Question
	var answerAll []*unity.Option
	var userResponse response.UserResponse
	err = utility.DB.QueryTable("exam_test").Filter("id", id).One(&examTest)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error015CantGetExamTest,
		})
	}
	err = utility.DB.QueryTable("user").Filter("id", IntUserId).One(&user)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error002ErrorQueryForGetAllUsers,
		})
	}
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.Role = user.Role
	userResponse.Email = user.Email
	userResponse.Phone = user.Phone
	userResponse.FullName = user.FullName
	if examTest.User.Id == IntUserId {
		_, err := utility.DB.QueryTable("question").Filter("exam_test_id", id).RelatedSel().All(&questionAll)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error016ErrorQueryForGetAllQuestions,
			})
		}
		examTest.Questions = questionAll
		var questionsResponse []response.QuestionResponse
		for i, question := range questionAll {
			questionsResponse = append(questionsResponse,
				response.QuestionResponse{
					Number:  question.Number,
					Content: question.Content,
					Answer:  question.Answer,
				},
			)
			_, err := utility.DB.QueryTable("option").Filter("question_id_id", question.Id).All(&answerAll)
			if err != nil {
				utility.FileLog.Println(err)
				return c.JSON(http.StatusInternalServerError, response.Message{
					Message: utility.Error017ErrorQueryForGetAllOptions,
				})
			}
			var OptionResponse []response.OptionResponse
			for _, answer := range answerAll {
				OptionResponse = append(OptionResponse, response.OptionResponse{
					Key:     answer.Key,
					Content: answer.Content,
				},
				)
			}
			questionsResponse[i].Options = OptionResponse
		}
		var customExamTestResponse = response.ExamTestResponse{
			Id:                examTest.Id,
			User:              userResponse,
			Date:              examTest.Date,
			Name:              examTest.Name,
			NumberOfQuestions: examTest.NumberOfQuestions,
			Subject:           examTest.Subject,
			Questions:         questionsResponse,
		}
		utility.FileLog.Println(user.Username + " get exam by id :" + examTest.Name)
		return c.JSON(http.StatusOK, customExamTestResponse)
	} else {
		return c.JSON(http.StatusUnauthorized, response.Message{
			Message: utility.Error018DontHavePermission,
		})
	}
}

func DownloadExam(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	QaId := c.Param("id")
	intQaId, err := strconv.ParseInt(QaId, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error061ExamIdInvalid,
		})
	}
	var examTest unity.ExamTest

	err = utility.DB.QueryTable("exam_test").Filter("id", intQaId).One(&examTest)

	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error015CantGetExamTest,
		})
	}
	utility.FileLog.Println(userName + " download " + examTest.Name)
	formatFile, err := formatDocxFileResult(examTest)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error045CantGetResult,
		})
	}
	return c.Attachment(formatFile, examTest.Name)
}

func DeleteExam(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)
	ExamId := c.Param("id")
	var examTest unity.ExamTest
	intExamId, err := strconv.ParseInt(ExamId, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error061ExamIdInvalid,
		})
	}
	err3 := utility.DB.QueryTable("exam_test").Filter("id", intExamId).One(&examTest)
	if err3 != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error015CantGetExamTest,
		})
	}
	if IntUserId == examTest.User.Id {
		_, err = utility.DB.QueryTable("exam_test").Filter("id", intExamId).Delete()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error062DeleteExamFailed,
			})
		}
		err2 := os.RemoveAll(examTest.Path)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error038RemoveFileError,
			})
		}
		message := response.Message{
			Message: "Delete exam successfully",
		}
		utility.FileLog.Println(userName + " delete file " + examTest.Name)
		return c.JSON(http.StatusOK, message)
	} else {
		return c.JSON(http.StatusUnauthorized, response.Message{
			Message: utility.Error018DontHavePermission,
		})
	}
}

var (
	specialCharacters = map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&apos;",
	}
)

func writeStringDocx(str string) string {
	output := ""
	for _, line := range strings.Split(str, "\n") {
		for key, value := range specialCharacters {
			line = strings.ReplaceAll(line, key, value)
		}
		output += "<w:t>" + line + "</w:t><w:br/>"
	}
	return output[:len(output)-7]
}

func formatDocxFileResult(exam unity.ExamTest) (string, error) {
	r, err := docx.ReadDocxFile("template/TestFormat.docx")
	var Questions []*unity.Question
	var OptionsAll []*unity.Option
	_, err = utility.DB.QueryTable("question").Filter("exam_test_id", exam.Id).All(&Questions)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	exam.Questions = Questions
	for _, question := range Questions {
		_, err := utility.DB.QueryTable("option").Filter("question_id_id", question.Id).All(&OptionsAll)
		if err != nil {
			utility.FileLog.Println(err)
			return err.Error(), err
		}
		question.Options = OptionsAll
	}

	table, err := ioutil.ReadFile("template/table.xml")
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	body, err := ioutil.ReadFile("template/body.xml")
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	tableContent := ""
	key := []string{"optionAContent", "optionBContent", "optionCContent", "optionDContent", "optionEContent", "optionFContent"}
	for num, question := range Questions {
		newtable := strings.ReplaceAll(string(table), "{{numberOfQuestion}}", writeStringDocx(strconv.Itoa(num+1)))
		newtable = strings.ReplaceAll(newtable, "{{QuestionContent}}", writeStringDocx(question.Content))
		for i := 0; i < 6; i++ {
			if i < len(question.Options) {
				newtable = strings.ReplaceAll(newtable, "{{"+key[i]+"}}", writeStringDocx(question.Options[i].Content))
			} else {
				newtable = strings.ReplaceAll(newtable, "{{"+key[i]+"}}", writeStringDocx(""))
			}
		}
		newtable = strings.ReplaceAll(newtable, "{{answers}}", writeStringDocx(question.Answer))
		tableContent += newtable
	}
	bodyContent := strings.ReplaceAll(string(body), "{{subject}}", writeStringDocx(exam.Subject))
	bodyContent = strings.ReplaceAll(bodyContent, "{{numberOfQuestions}}", writeStringDocx(strconv.Itoa(int(exam.NumberOfQuestions))))
	bodyContent = strings.ReplaceAll(bodyContent, "{{table}}", tableContent)
	editFile := r.Editable()
	editFile.SetContent(bodyContent)
	extension := filepath.Ext(exam.Path)
	newFormatFile := strings.ReplaceAll(exam.Path, extension, "-"+strconv.Itoa(int(exam.Id))+".docx")
	f, err := os.Create(newFormatFile)
	err = editFile.WriteToFile(f.Name())
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	return newFormatFile, nil
}
