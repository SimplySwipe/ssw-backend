package handlers

import (
	"SimplySwipe/models"
	"SimplySwipe/utils"
	"context"
	"os"
	"time"

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
	audience := os.Getenv("JWT_AUDIENCE")
	issuer := os.Getenv("JWT_ISSUER")
	role := "guest"

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	userID := payload.Claims["sub"].(string)

	accessToken, err := utils.GenerateJWT(userID, email, role, audience, issuer, time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}
	refreshToken, err := utils.GenerateJWT(userID, email, role, audience, issuer, 24*7*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}
	c.JSON(200, gin.H{
		"accesToken":   accessToken,
		"refreshToken": refreshToken,
		"user": gin.H{
			"id":    userID,
			"email": email,
			"name":  name,
		},
	})

}

func TestToken(c *gin.Context) {

	audience := os.Getenv("JWT_AUDIENCE")
	issuer := os.Getenv("JWT_ISSUER")
	role := "guest"

	userID := "5"
	email := "tian.istenic@gmail.com"
	name := "tian"
	accessToken, err := utils.GenerateJWT(userID, email, role, audience, issuer, time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}
	c.JSON(200, gin.H{
		"accesToken":   accessToken,
		"refreshToken": accessToken,
		"user": gin.H{
			"id":    userID,
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
