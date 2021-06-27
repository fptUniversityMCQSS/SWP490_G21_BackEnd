package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func History(c echo.Context) error {
	return c.File("views/history.html")
}
func GetExamById(c echo.Context) error {
	token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	values, _ := jwt.Parse(token, nil)
	claims := values.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	o := orm.NewOrm()
	var examTest model.ExamTest
	var questionAll []*model.Question
	var answerAll []*model.Option
	o.QueryTable("exam_test").Filter("id", id).All(&examTest)
	if examTest.User.Id == IntUserId {
		o.QueryTable("question").Filter("exam_test_id", id).RelatedSel().All(&questionAll)
		examTest.Questions = questionAll

		var questionsResponse []response.QuestionResponse

		for _, question := range questionAll {
			questionsResponse = append(questionsResponse,
				response.QuestionResponse{
					Number: question.Number,
					Content: question.
				}
			)

			o.QueryTable("option").Filter("question_id_id", question.Id).All(&answerAll)
			question.Options = answerAll
		}

		var customExamTestResponse = response.ExamTestResponse{
			Id: examTest.Id,
			Date: examTest.Date,
			Name: examTest.Name,
			Questions: []response.QuestionResponse{
				response.QuestionResponse{

				},
			},
		}

		return c.JSON(http.StatusOK, examTest)
	} else {
		return c.JSON(http.StatusUnauthorized, "You dont have permission to access this")
	}
	return nil

}
