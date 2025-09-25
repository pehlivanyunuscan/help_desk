package handlers

import "github.com/gofiber/fiber/v2"

func CreateFaultReport(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Fault report created",
	})
}
