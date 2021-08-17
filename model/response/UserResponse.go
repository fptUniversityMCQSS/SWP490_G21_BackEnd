package response

type UserResponse struct {
	Id       int64  `orm:"pk;auto" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	FullName string `json:"full name"`
	Role     string `json:"role" form:"role"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}
