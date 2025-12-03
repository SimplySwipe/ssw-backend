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
	refreshToken, err := db.GenerateRefreshToken()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate the refresh token"})
		return
	}
	_, err = db.InsertRefreshToken(c.Request.Context(), db.Pool, user.ID, refreshToken, time.Now().Add(24*7*time.Hour))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to store the refresh token"})
		return
	}
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", "", true, true)

	c.JSON(200, gin.H{
		"accesToken": accessToken,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})

}

func RefreshToken(c *gin.Context) {
	cookieToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "refresh token missing in cookie"})
		return
	}
	refreshToken, err := db.GetRefreshToken(c.Request.Context(), db.Pool, cookieToken)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	if refreshToken.Token != cookieToken {
		c.JSON(401, gin.H{"error": "refresh token missing or invalid"})
		return
	}
	if refreshToken.Revoked {
		c.JSON(401, gin.H{"error": "refresh token revoked"})
		return
	}
	if refreshToken.Used {
		c.JSON(401, gin.H{"error": "refresh token used"})
		return
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		c.JSON(401, gin.H{"error": "refresh token expired"})
		return
	}
	// generate new token
	// mark old as used/revoked/expired
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logout"})
}
