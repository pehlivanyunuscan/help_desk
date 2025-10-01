package handlers

import (
	"help_desk/database"
	"help_desk/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateFaultReport(c *fiber.Ctx) error {
	var input models.CreateFaultReportRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if input.Title == "" || input.MachineID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	// clock alanını time.Time türüne dönüştürür
	timestamp := time.Unix(input.Clock, 0)

	// JWT ile gelen kullanıcı bilgilerini al
	userIface := c.Locals("user")
	username := ""
	if user, ok := userIface.(models.User); ok {
		username = user.Username
	}

	report := models.FaultReport{
		Title:           input.Title,
		UserDescription: input.UserDescription,
		Timestamp:       timestamp,
		MachineID:       input.MachineID,
		ReportedBy:      username, // JWT'den alınan kullanıcı adı
		Asset:           input.Asset,
	}
	if err := database.DB.Create(&report).Error; err != nil {
		log.Printf("Failed to create fault report for user %s: %v", username, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create fault report",
		})
	}

	// Başarılı report oluşturma log kaydı
	log.Printf("User %s created fault report (ID: %d, Title: %s, Machine: %s) at %s",
		username, report.ID, report.Title, report.MachineID, time.Now().Format("2006-01-02 15:04:05"))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Fault report created successfully",
		"report": fiber.Map{
			"id":               report.ID,
			"title":            report.Title,
			"asset":            report.Asset,
			"machine_id":       report.MachineID,
			"priority":         report.Priority,
			"user_description": report.UserDescription,
			"timestamp":        report.Timestamp,
			"reported_by":      report.ReportedBy,
		},
	})
}

func GetFaultReports(c *fiber.Ctx) error {
	// JWT ile gelen kullanıcı bilgilerini al
	userIface := c.Locals("user")
	if user, ok := userIface.(models.User); ok {
		// Log kaydı
		log.Printf("User %s accessed fault reports", user.Username)
	}

	var reports []models.FaultReport
	if err := database.DB.Find(&reports).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve fault reports",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Fault reports retrieved successfully",
		"data":    reports,
		"count":   len(reports),
	})
}

func GetFaultReportByID(c *fiber.Ctx) error {
	// JWT ile gelen kullanıcı bilgilerini al
	userIface := c.Locals("user")
	if user, ok := userIface.(models.User); ok {
		// Log kaydı
		log.Printf("User %s accessed fault report with ID %s", user.Username, c.Params("id"))
	}

	id := c.Params("id")
	var report models.FaultReport
	if err := database.DB.First(&report, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Fault report not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Fault report retrieved successfully",
		"data":    report,
	})
}
