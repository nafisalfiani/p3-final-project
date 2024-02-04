package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/header"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
)

type contextKey string

const (
	userAuthInfo contextKey = "UserAuthInfo"
)

type Interface interface {
	VerifyToken(ctx context.Context, tokenId string) (User, error)
	CreateToken(ctx context.Context, user User) (Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (RefreshTokenResponse, error)

	SetUserAuthInfo(ctx context.Context, user User, token *Token) context.Context
	GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error)
}

type auth struct {
	log        log.Interface
	json       parser.JSONInterface
	httpClient *http.Client
	conf       Config
}

type Config struct {
	AccessTokenKey       string        `env:"AUTH_ACCESS_TOKEN_KEY"`
	AccessTokenDuration  time.Duration `env:"AUTH_ACCESS_TOKEN_DURATION"`
	RefreshTokenKey      string        `env:"AUTH_REFRESH_TOKEN_KEY"`
	RefreshTokenDuration time.Duration `env:"AUTH_REFRESH_TOKEN_DURATION"`
}

func Init(cfg Config, log log.Interface, json parser.JSONInterface, httpClient *http.Client) Interface {
	return &auth{
		log:        log,
		json:       json,
		httpClient: httpClient,
		conf:       cfg,
	}
}

func (a *auth) VerifyToken(ctx context.Context, accessToken string) (User, error) {
	var user User
	claims := jwt.MapClaims{}
	refreshTokenParser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	_, err := refreshTokenParser.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.conf.AccessTokenKey), nil
	})
	if err != nil {
		return user, err
	}

	user, expiry, err := a.extractClaims(claims)
	if err != nil {
		return user, err
	}

	if time.Unix(expiry, 0).Before(time.Now()) {
		return user, errors.NewWithCode(codes.CodeAuthAccessTokenExpired, "access token expired")
	}

	return user, nil
}

func (a *auth) CreateToken(ctx context.Context, user User) (Token, error) {
	accessExpiryTime := time.Now().Add(a.conf.AccessTokenDuration).Unix()
	accessClaims := jwt.MapClaims{
		"user:id":       user.Id,
		"user:name":     user.Name,
		"user:email":    user.Email,
		"user:username": user.Username,
		"user:role":     user.Role,
		"user:scopes":   user.Scopes,
		"expiry":        accessExpiryTime,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(a.conf.AccessTokenKey))
	if err != nil {
		return Token{}, err
	}

	refreshExpiryTime := time.Now().Add(a.conf.RefreshTokenDuration).Unix()
	refreshClaims := jwt.MapClaims{
		"user:id":       user.Id,
		"user:name":     user.Name,
		"user:email":    user.Email,
		"user:username": user.Username,
		"user:role":     user.Role,
		"user:scopes":   user.Scopes,
		"expiry":        refreshExpiryTime,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(a.conf.RefreshTokenKey))
	if err != nil {
		return Token{}, err
	}

	token := Token{
		TokenType:        header.AuthorizationBearer,
		AccessToken:      accessTokenString,
		AccessExpiresIn:  accessExpiryTime,
		RefreshToken:     refreshTokenString,
		RefreshExpiresIn: refreshExpiryTime,
	}

	return token, nil
}

func (a *auth) RefreshToken(ctx context.Context, refreshToken string) (RefreshTokenResponse, error) {
	// Verify and decode the refresh token.
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenParser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	_, err := refreshTokenParser.ParseWithClaims(refreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.conf.RefreshTokenKey), nil
	})
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	user, expiry, err := a.extractClaims(refreshTokenClaims)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	if time.Unix(expiry, 0).Before(time.Now()) {
		return RefreshTokenResponse{}, errors.NewWithCode(codes.CodeAuthRefreshTokenExpired, "refresh token expired")
	}

	// Generate a new access token with updated claims.
	accessExpiryTime := time.Now().Add(a.conf.AccessTokenDuration).Unix()
	accessClaims := jwt.MapClaims{
		"user:id":       user.Id,
		"user:name":     user.Name,
		"user:email":    user.Email,
		"user:username": user.Username,
		"user:role":     user.Role,
		"user:scopes":   user.Scopes,
		"expiry":        accessExpiryTime,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(a.conf.AccessTokenKey))
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	// Return the new access & refresh token and its expiration time.
	response := RefreshTokenResponse{
		TokenType:        header.AuthorizationBearer,
		AccessToken:      accessTokenString,
		ExpiresIn:        accessExpiryTime,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: time.Now().Add(a.conf.RefreshTokenDuration).Unix(),
		UserID:           user.Id,
	}

	return response, nil
}

func (a *auth) extractClaims(claims jwt.MapClaims) (User, int64, error) {
	var user User
	var expiredIn int64
	var ok bool

	user.Id, ok = claims["user:id"].(string)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid id format")
	}

	user.Name, ok = claims["user:name"].(string)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid name format")
	}

	user.Email, ok = claims["user:email"].(string)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid email format")
	}

	user.Username, ok = claims["user:username"].(string)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid username format")
	}

	user.Role, ok = claims["user:role"].(string)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid role format")
	}

	scopes, ok := claims["user:scopes"].([]interface{})
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "Invalid scopes format")
	}

	for i := range scopes {
		if s, ok := scopes[i].(string); ok {
			user.Scopes = append(user.Scopes, s)
		}
	}

	expiry, ok := claims["expiry"].(float64)
	if !ok {
		return user, expiredIn, errors.NewWithCode(codes.CodeAuth, "invalid expiry format")
	}

	return user, int64(expiry), nil
}

func (a *auth) SetUserAuthInfo(ctx context.Context, user User, token *Token) context.Context {
	userauth := UserAuthInfo{
		User:  user,
		Token: token,
	}

	return context.WithValue(ctx, userAuthInfo, userauth)
}

func (a *auth) GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error) {
	user, ok := ctx.Value(userAuthInfo).(UserAuthInfo)
	if !ok {
		return user, errors.NewWithCode(codes.CodeAuthFailure, "failed getting user auth info")
	}

	return user, nil
}
