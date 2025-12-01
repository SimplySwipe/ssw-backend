package middleware

import (
	"SimplySwipe/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}
		tokenString := authHeader[7:]
		var claims models.Claims
		secret := os.Getenv("JWT_SECRET")
		audience := os.Getenv("JWT_AUDIENCE")
		issuer := os.Getenv("JWT_ISSUER")
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}
		if claims.Audience != audience {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token audience"})
			return
		}
		if claims.Issuer != issuer {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token issuer"})
			return
		}
		c.Set("userClaims", claims)
		c.Next()
	}
}
