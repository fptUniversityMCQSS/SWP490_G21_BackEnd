package controller

import "github.com/labstack/echo/v4"

func Knowledge(c echo.Context) error {
	return c.File("views/knowledge.html")
}

func KnowledgeUpload(c echo.Context) error {
	return c.File("views/knowledgeUpload.html")
}
