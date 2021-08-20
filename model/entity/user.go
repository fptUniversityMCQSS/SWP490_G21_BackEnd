package entity

type User struct {
	Id        int64        `orm:"pk;auto" json:"id" form:"id"`
	Username  string       `json:"username" form:"username"`
	Password  string       `json:"password" form:"password"`
	Email     string       `json:"email" form:"email"`
	Phone     string       `json:"phone" form:"phone"`
	Role      string       `json:"role" form:"role"`
	FullName  string       `json:"fullName" form:"fullName"`
	Knowledge []*Knowledge `orm:"reverse(many)" json:"list_knowledge" form:"list_knowledge"`
}
