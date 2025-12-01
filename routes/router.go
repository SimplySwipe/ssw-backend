package routes

import (
	"SimplySwipe/handlers"
	"SimplySwipe/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RateLimitMiddleware())

	{
		api := router.Group("/api")
		api.GET("/ping", handlers.Ping)

		auth := api.Group("/auth")
		auth.POST("/oauth/google", handlers.GoogleOAuth)
		auth.POST("/test-token", handlers.TestToken)

		auth.POST("/refresh", handlers.RefreshToken)
		auth.POST("/logout", handlers.Logout)

		user := api.Group("/user")
		user.Use(middleware.JWTAuth())
		user.GET("/profile", handlers.GetUserProfile)
		user.PUT("/profile", handlers.UpdateUserProfile)

		jobs := api.Group("/jobs")
		jobs.Use(middleware.JWTAuth())

		internal := api.Group("/internal")
		internal.POST("/scraper/push", handlers.ScraperPush)
	}

	return router
}
