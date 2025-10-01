package main

import (
	"help_desk/database"
	"help_desk/handlers"
	"help_desk/middleware"
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

	app.Post("/login", handlers.Login)

	app.Post("/send_helpdesk", middleware.RequireAuth, handlers.CreateFaultReport)
	app.Get("/get_faultreports", middleware.RequireAuth, handlers.GetFaultReports)
	app.Get("/get_faultreport/:id", middleware.RequireAuth, handlers.GetFaultReportByID)

	app.Listen(":3000")
}
