package handlers

import (
	"SimplySwipe/models"

	"github.com/gin-gonic/gin"
)

func ScraperPush(c *gin.Context) {
	var input models.ScraperJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	c.JSON(200, gin.H{"recieved": input})
}
