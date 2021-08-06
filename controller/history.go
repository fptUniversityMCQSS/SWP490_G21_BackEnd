package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"github.com/astaxie/beego/orm"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func History(c echo.Context) error {
	o := orm.NewOrm()
	var history []*model.ExamTest
	var hist []*response.HistoryResponse

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	//log.Println("test: " + userid.(string))

	_, err := o.QueryTable("exam_test").Filter("user_id", userId).All(&history)

	//if has problem in connection
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}

	//add selected data to knowledge_Res list
	for _, h := range history {
		var his = new(response.HistoryResponse)
		his.Id = h.Id
		his.Name = h.Name
		his.Date = h.Date
		hist = append(hist, his)
	}
	log.Printf(userName + " get list history")
	return c.JSON(http.StatusOK, hist)

}
func GetExamById(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	o := orm.NewOrm()
	var examTest model.ExamTest
	var user model.User
	var questionAll []*model.Question
	var answerAll []*model.Option
	var userResponse response.UserResponse
	o.QueryTable("exam_test").Filter("id", id).All(&examTest)
	o.QueryTable("user").Filter("id", IntUserId).All(&user)
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.Role = user.Role
	if examTest.User.Id == IntUserId {
		o.QueryTable("question").Filter("exam_test_id", id).RelatedSel().All(&questionAll)
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
		log.Printf(user.Username + " get exam by id :" + examTest.Name)
		return c.JSON(http.StatusOK, customExamTestResponse)
	} else {
		return c.JSON(http.StatusUnauthorized, "You dont have permission to access this")
	}
	return nil

}

func DownloadExam(c echo.Context) error {
	o := orm.NewOrm()
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	QaId := c.Param("id")
	intQaId, _ := strconv.ParseInt(QaId, 10, 64)
	var examTest model.ExamTest

	err := o.QueryTable("exam_test").Filter("id", intQaId).One(&examTest)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
		return err
	}
	log.Printf(userName + " download " + examTest.Name)
	return c.Attachment("examtest/"+examTest.Name, examTest.Name)
}
