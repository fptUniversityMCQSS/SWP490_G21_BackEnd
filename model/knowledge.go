package model

import "time"

type Knowledge struct {
	Id     int64     `orm:"pk;auto"`
	Name   string    `json:"knowledgeName" form:"knowledgeName"`
	Date   time.Time `orm:"auto_now_add;type(datetime)" json:"knowledgeDate" form:"knowledgeDate"`
	User   *User     `orm:"rel(fk);on_delete(cascade)"`
	Path   string
	Status string
}
