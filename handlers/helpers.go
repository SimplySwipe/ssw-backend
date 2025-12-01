package handlers

import (
	"SimplySwipe/models"

	"github.com/gin-gonic/gin"
)

func GetUserClaims(c *gin.Context) (models.Claims, bool) {
	claims, exist := c.Get("userClaims")
	if !exist {
		return models.Claims{}, false
	}
	userClaims, ok := claims.(models.Claims)
	return userClaims, ok
}
