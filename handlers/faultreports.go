package handlers

import (
	"help_desk/database"
	"help_desk/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Gelen istek için kullanılacak yapı
type CreateFaultReportRequest struct {
	Title           string `json:"title"`
	UserDescription string `json:"user_description"`
	Clock           int64  `json:"clock"`
	MachineID       string `json:"machine_id"`
	Asset           string `json:"asset"`
}

func CreateFaultReport(c *fiber.Ctx) error {
	var input CreateFaultReportRequest
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

	report := models.FaultReport{
		Title:           input.Title,
		UserDescription: input.UserDescription,
		Timestamp:       timestamp,
		MachineID:       input.MachineID,
	}
	if err := database.DB.Create(&report).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create fault report",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(report)
}
