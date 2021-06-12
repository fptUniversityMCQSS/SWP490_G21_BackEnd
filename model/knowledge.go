package model

import "time"

type Knowledge struct {
	Id int64         `orm:"pk"`
	Name string
	Date time.Time
	User *User       `orm:"rel(fk)"`
}
