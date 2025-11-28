package main

import (
	"SimplySwipe/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
