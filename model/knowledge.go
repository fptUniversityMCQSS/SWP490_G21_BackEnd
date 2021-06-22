package model

import "time"

type Knowledge struct {
	Id   int64     `orm:"pk"`
	Name string    `json:"knowledgeName" form:"knowledgeName"`
	Date time.Time `orm:"auto_now_add" json:"knowledgeDate" form:"knowledgeDate"`
	User *User     `orm:"rel(fk)"`
}