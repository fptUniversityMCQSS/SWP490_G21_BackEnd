package model

type User struct {
	Id int64         `orm:"pk"`
	Username string
	Password string
	Role     string
	Knowledge []*Knowledge `orm:"reverse(many)"`
	Question []*Question   `orm:"reverse(many)"`
}
