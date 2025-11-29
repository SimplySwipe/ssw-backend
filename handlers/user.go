package handlers

import (
	"SimplySwipe/db"
	"SimplySwipe/models"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	claims, exist := c.Get("userClaims")
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		return
	}

	userClaims, ok := claims.(models.Claims)
	if !ok {
		c.JSON(401, gin.H{"error": "No claims"})
		return
	}
	userID := userClaims.UserID
	user, err := db.GetUserByID(c.Request.Context(), db.Pool, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)

}

func UpdateUserProfile(c *gin.Context) {

}
