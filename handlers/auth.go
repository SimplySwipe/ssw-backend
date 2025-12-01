package handlers

import (
	"SimplySwipe/db"
	"SimplySwipe/models"
	"SimplySwipe/utils"
	"context"
	"log"
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
	if payload.Audience != googleClientID {
		c.JSON(401, gin.H{"error": "Invalid token audience"})
		return
	}
	if payload.Issuer != "https://accounts.google.com" && payload.Issuer != "accounts.google.com" {
		c.JSON(401, gin.H{"error": "Invalid token issuer"})
		return
	}

	emailVerified := false
	if v, ok := payload.Claims["email_verified"]; ok {
		if b, ok := v.(bool); ok {
			emailVerified = b
		}
	}
	if !emailVerified {
		c.JSON(401, gin.H{"error": "Google account not verified"})
		return
	}

	audience := os.Getenv("JWT_AUDIENCE")
	issuer := os.Getenv("JWT_ISSUER")
	role := "guest"

	email := ""
	if v, ok := payload.Claims["email"]; ok {
		email, _ = v.(string)
	}
	name := ""
	if v, ok := payload.Claims["name"]; ok {
		name, _ = v.(string)
	}
	photoURL := ""
	if v, ok := payload.Claims["picture"]; ok {
		photoURL, _ = v.(string)
	}
	googleID := payload.Subject
	// email := payload.Claims["email"].(string)
	// name := payload.Claims["name"].(string)
	// photoURL := payload.Claims["picture"].(string)
	log.Printf("%#v", payload.Claims)

	user, err := db.GetOrCreateUserByGoogleID(
		c.Request.Context(),
		db.Pool,
		googleID,
		email,
		name,
		photoURL,
	)
	if err != nil {
		log.Println("error, failed to get or create user", err)
		c.JSON(500, gin.H{"error": "Failed to get or creat user"})
	}
	if user == nil {
		c.JSON(500, gin.H{"error": "User not found or created"})
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, email, role, audience, issuer, time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}
	refreshToken, err := utils.GenerateJWT(user.ID, email, role, audience, issuer, 24*7*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}
	c.JSON(200, gin.H{
		"accesToken":   accessToken,
		"refreshToken": refreshToken,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})

}

func RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{"message": "RefreshToken"})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logout"})
}
