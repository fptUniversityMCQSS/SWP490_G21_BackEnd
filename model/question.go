package model

import "time"

type Question struct {
	Id int64  `orm:"pk"`
	Content string
	Date    time.Time
	Options []*Option  `orm:"reverse(many)"`
	Answer  *Option `orm:"null;rel(one);on_delete(set_null)"`
	User    *User    `orm:"rel(fk)"`
}
