package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
