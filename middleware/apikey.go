package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedKey := os.Getenv("INTERNAL_API_KEY")
		providedKey := c.GetHeader("X-API-Key")
		if providedKey != expectedKey || providedKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or missing api key"})
			return
		}
		c.Next()
	}
}
