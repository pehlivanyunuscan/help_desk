package main

import (
	"help_desk/database"
	"help_desk/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	app := fiber.New()

	app.Post("/faultreports", handlers.CreateFaultReport)

	app.Listen(":3000")
}
