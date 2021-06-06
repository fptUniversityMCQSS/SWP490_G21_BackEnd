package model

import "time"

type Question struct {
	Content string
	Date    time.Time
	Options []Option
	Answer  Option
	User    User
}
