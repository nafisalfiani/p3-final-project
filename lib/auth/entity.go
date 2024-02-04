package auth

type User struct {
	Id       string   `db:"id" json:"id" swaggerignore:"true"`
	Name     string   `db:"name" json:"name"`
	Username string   `db:"username" json:"username"`
	Email    string   `db:"email" json:"email"`
	Role     string   `db:"role" json:"role"`
	Scopes   []string `db:"-" json:"scopes"`
}

type Token struct {
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	AccessExpiresIn  int64  `json:"access_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int64  `json:"-"`
}

type UserAuthInfo struct {
	User  User   `json:"user"`
	Token *Token `json:"token"`
}

type RefreshTokenResponse struct {
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int64  `json:"-"`
	UserID           string `json:"user_id"`
}
