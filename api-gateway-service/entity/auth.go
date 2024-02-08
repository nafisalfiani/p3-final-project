package entity

type RegisterRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type VerifyEmailRequest struct {
	Id string `uri:"id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	AccessExpiresIn int64  `json:"access_expires_in"`
}
