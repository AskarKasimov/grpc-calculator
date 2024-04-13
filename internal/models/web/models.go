package web

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	UserId int64 `json:"id"`
	jwt.RegisteredClaims
}
