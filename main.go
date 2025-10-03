package main

import (
	"help_desk/database"
	"help_desk/handlers"
	"help_desk/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "help_desk/docs" // swagger docs
)

// @title Help Desk API
// @version 1.0
// @description Help Desk API for fault reporting system
// @contact.name API Support
// @contact.email support@helpdesk.com
// @host 10.67.67.22:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your JWT token in the format: Bearer <your_token_here>

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	app := fiber.New()

	// Swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/login", handlers.Login)

	app.Post("/fault-reports", middleware.RequireAuth, handlers.CreateFaultReport)
	app.Get("/fault-reports", middleware.RequireAuth, handlers.GetFaultReports)
	app.Get("/fault-reports/:id", middleware.RequireAuth, handlers.GetFaultReportByID)

	app.Listen("0.0.0.0:3000")
}
