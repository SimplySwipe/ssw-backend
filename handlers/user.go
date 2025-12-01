package handlers

import (
	"SimplySwipe/db"
	"SimplySwipe/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userClaims, ok := GetUserClaims(c)
	if !ok {
		c.JSON(401, gin.H{"error": "unauthorized", "message": "Missing or invalid authentication claims"})
		return
	}
	userID := userClaims.UserID
	user, err := db.GetUserByID(c.Request.Context(), db.Pool, userID)
	if err != nil {
		log.Printf("[GetUserProfile] DB error: %v", err)
		c.JSON(500, gin.H{"error": "internal_server_error", "message": "Could not retrieve user profile"})
		return
	}
	if user == nil {
		c.JSON(404, gin.H{"error": "not_found", "message": "User not found"})
		return
	}
	c.JSON(200, user)

}

func UpdateUserProfile(c *gin.Context) {
	userClaims, ok := GetUserClaims(c)
	if !ok {
		c.JSON(401, gin.H{"error": "unauthorized", "message": "Missing or invalid authentication claims"})
		return
	}
	userID := userClaims.UserID

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "bad_request", "message": "Invalid input: " + err.Error()})
		return
	}

	updateUser, err := db.UpdateUser(
		c.Request.Context(),
		db.Pool,
		userID,
		req.Name,
		req.Phone,
		req.PhotoURL,
	)
	if err != nil {
		log.Printf("[UpdateUserProfile] DB error: %v", err)
		c.JSON(500, gin.H{"error": "internal_server_error", "message": "Failed to update user profile"})
		return
	}
	if updateUser == nil {
		c.JSON(404, gin.H{"error": "not_found", "message": "User not found"})
		return
	}
	c.JSON(200, updateUser)
}
