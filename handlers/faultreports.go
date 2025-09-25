package handlers

import (
	"help_desk/database"
	"help_desk/models"
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

	report := models.FaultReport{
		Title:           input.Title,
		UserDescription: input.UserDescription,
		Timestamp:       timestamp,
		MachineID:       input.MachineID,
		ReportedBy:      "system", // Varsayılan olarak "system" atandı, gerçek kullanıcı bilgisi eklenebilir
		Asset:           input.Asset,
	}
	if err := database.DB.Create(&report).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create fault report",
		})
	}

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
	var reports []models.FaultReport
	if err := database.DB.Find(&reports).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve fault reports",
		})
	}
	return c.JSON(reports)
}

func GetFaultReportByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var report models.FaultReport
	if err := database.DB.First(&report, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Fault report not found",
		})
	}
	return c.JSON(report)
}
