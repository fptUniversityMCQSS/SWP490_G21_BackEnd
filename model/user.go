package model

type User struct {
	Id        int64        `orm:"pk" json:"id" form:"id"`
	Username  string       `json:"username" form:"username"`
	Password  string       `json:"password" form:"password"`
	Role      string       `json:"role" form:"role"`
	Knowledge []*Knowledge `orm:"reverse(many)" json:"list_knowledge" form:"list_knowledge"`
	Question  []*Question  `orm:"reverse(many)"   json:"list_question" form:"list_question"`
}
