package model

import "time"

type Question struct {
	Id      int64     `orm:"pk"`
	Content string    `json:"question_content" form:"question_content"`
	Date    time.Time `orm:"auto_now_add" json:"question_date" form:"question_date"`
	Options []*Option `orm:"reverse(many)"`
	Answer  *Option   `orm:"null;rel(one);on_delete(set_null)" json:"answer" form:"answer"`
	User    *User     `orm:"rel(fk)"`
}
