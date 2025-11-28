package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID   string `json:"sub"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Audience string `json:"audience"`
	Issuer   string `json:"issuer"`
	jwt.RegisteredClaims
}
