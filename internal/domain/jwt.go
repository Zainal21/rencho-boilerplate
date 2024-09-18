package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtAuthUseCase interface {
	CreateUser(email, password string) (authUID string, err error)
	VerifyToken(token string) (authUID string, err error)
	GetAccessToken(email, password string) (accessToken string, err error)
}

type JWTAccessTokenClaims struct {
	UserUID string `json:"user_uid"`
	jwt.RegisteredClaims
}

type JWTUtil interface {
	GenerateAccessToken(userUID string) (string, time.Time, error)
	GenerateRefreshToken(userUID string) (string, time.Time, error)
	ParseUserUID(tokenString string, isAccessToken bool) (string, error)
	Refresh(refreshToken string) (string, time.Time, error)
}
