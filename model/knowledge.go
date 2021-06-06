package model

import "time"

type Knowledge struct {
	Name string
	Date time.Time
	User User
}
