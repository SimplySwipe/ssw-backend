package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	{
		api := router.Group("/api")
		api.GET("/ping", Ping)

		auth := api.Group("/auth")
		auth.GET("/ping", Ping)

		user := api.Group("/me")
		user.GET("/ping", Ping)

		jobs := api.Group("/jobs")
		jobs.GET("/ping", Ping)

		internal := api.Group("/internal")
		internal.GET("/ping", Ping)
	}

	return router
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
