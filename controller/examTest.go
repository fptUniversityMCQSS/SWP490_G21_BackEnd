package controller

import (
	"SWP490_G21_Backend/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DownloadExam(c echo.Context) error {
	o := orm.NewOrm()
	QaId := c.Param("id")
	intQaId, _ := strconv.ParseInt(QaId, 10, 64)
	var examTest model.ExamTest

	err := o.QueryTable("exam_test").Filter("id", intQaId).One(&examTest)

	//if has problem in connection
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	return c.Attachment("examtest/"+examTest.Name, examTest.Name)
}
