package handlers

import (
	"SimplySwipe/models"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func GoogleOAuth(c *gin.Context) {
	var req models.GoogleOAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")

	payload, err := idtoken.Validate(context.Background(), req.IDToken, googleClientID)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid Google ID token"})
		return
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	sub := payload.Claims["sub"].(string)

	c.JSON(200, gin.H{
		"accesToken":   "nothing-yet",
		"refreshToken": "nothing-yet",
		"user": gin.H{
			"id":    sub,
			"email": email,
			"name":  name,
		},
	})
}

func RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{"message": "RefreshToken"})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logout"})
}
