package model

import "time"

type ExamTest struct {
	Id        int64       `orm:"pk;auto"`
	Questions []*Question `orm:"reverse(many)"`
	Date      time.Time   `orm:"auto_now_add" json:"test_date" form:"test_date"`
	User      *User       `orm:"rel(fk)"`
}
