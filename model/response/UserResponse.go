package response

type UserResponse struct {
	Id       int64  `orm:"pk;auto" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`
}
