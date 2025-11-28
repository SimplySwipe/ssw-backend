package utils

import (
	"SimplySwipe/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, email, role, audience, issuer string, duration time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := models.Claims{
		UserID:   userID,
		Email:    email,
		Role:     role,
		Audience: audience,
		Issuer:   issuer,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
			Audience:  []string{audience},
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}
