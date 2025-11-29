package main

import (
	"SimplySwipe/db"
	"SimplySwipe/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	router := routes.SetupRouter()
	router.Run(":8080")
}
