package model

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}
