package response

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	IsAdmin  bool   `json:"is_admin"`
	JWTToken string `json:"jwt_token,omitempty"`
}
