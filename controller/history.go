package controller

import (
	"SWP490_G21_Backend/model"
	"SWP490_G21_Backend/model/response"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
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
