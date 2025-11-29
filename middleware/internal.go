package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedKey := os.Getenv("INTERNAL_API_KEY")
		receivedKey := c.GetHeader("X-API-Key")
		if receivedKey == "" || receivedKey != expectedKey {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or missing API key"})
			return
		}
		c.Next()
	}
}
