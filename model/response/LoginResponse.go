package response

type LoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}
