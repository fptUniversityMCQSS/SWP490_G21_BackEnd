package response

type LoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
