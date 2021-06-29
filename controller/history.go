package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func History(c echo.Context) error {
	o := orm.NewOrm()
	var history []*model.ExamTest
	var hist []*response.HistoryResponse

	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userid := claims["userId"]
	//log.Println("test: " + userid.(string))

	qs, err := o.QueryTable("exam_test").Filter("user_id", userid).All(&history)

	//if has problem in connection
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}

	//add selected data to knowledge_Res list
	for _, h := range history {
		var his = new(response.HistoryResponse)
		his.Id = h.Id
		his.Date = h.Date
		hist = append(hist, his)
	}
	fmt.Printf("%d history read \n", qs)
	return c.JSON(http.StatusOK, hist)

}
func GetExamById(c echo.Context) error {
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	o := orm.NewOrm()
	var examTest model.ExamTest
	var user model.User
	var questionAll []*model.Question
	var answerAll []*model.Option
	var userResponse response.UserResponse
	var answerResponse response.OptionResponse
	o.QueryTable("exam_test").Filter("id", id).All(&examTest)
	o.QueryTable("user").Filter("id", IntUserId).All(&user)
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	if examTest.User.Id == IntUserId {
		o.QueryTable("question").Filter("exam_test_id", id).RelatedSel().All(&questionAll)
		examTest.Questions = questionAll

		var questionsResponse []response.QuestionResponse
		for i, question := range questionAll {
			answerResponse.Key = question.Answer.Key
			answerResponse.Content = question.Answer.Content
			questionsResponse = append(questionsResponse,
				response.QuestionResponse{
					Number:  question.Number,
					Content: question.Content,
					Answer:  answerResponse,
				},
			)
			o.QueryTable("option").Filter("question_id_id", question.Id).All(&answerAll)
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
			Id:        examTest.Id,
			User:      userResponse,
			Date:      examTest.Date,
			Name:      examTest.Name,
			Questions: questionsResponse,
		}
		return c.JSON(http.StatusOK, customExamTestResponse)
	} else {
		return c.JSON(http.StatusUnauthorized, "You dont have permission to access this")
	}
	return nil

}
