package controller

import (
	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
	"lib/model"
	"net/http"
)

func Knowledge(c echo.Context) error {
	o := orm.NewOrm()
	var posts []*model.Knowledge
	o.QueryTable("knowledge").Filter("user_id", "%").RelatedSel().All(&posts)

	return c.JSON(http.StatusOK, posts)
}

func KnowledgeUpload(c echo.Context) error {
	return c.File("views/knowledgeUpload.html")
}
