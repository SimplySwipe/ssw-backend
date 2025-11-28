package handlers

import "github.com/gin-gonic/gin"

func GoogleOAuth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GoogleOauth"})
}

func RefreshToken(c *gin.Context) {
	c.JSON(200, gin.H{"message": "RefreshToken"})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logout"})
}
