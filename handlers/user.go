package handlers

import (
	"SimplySwipe/models"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	claims, exist := c.Get("userClaims")
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
	}

	userClaims, ok := claims.(models.Claims)
	if !ok {
		c.JSON(401, gin.H{"error": "no claims"})
	}
	userID := userClaims.UserID
}

func UpdateUserProfile(c *gin.Context) {

}
