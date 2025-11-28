package routes

import (
	"SimplySwipe/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	{
		api := router.Group("/api")
		api.GET("/ping", handlers.Ping)

		auth := api.Group("/auth")
		auth.GET("/ping", handlers.Ping)

		user := api.Group("/me")
		user.GET("/ping", handlers.Ping)

		jobs := api.Group("/jobs")
		jobs.GET("/ping", handlers.Ping)

		internal := api.Group("/internal")
		internal.GET("/ping", handlers.Ping)
	}

	return router
}
