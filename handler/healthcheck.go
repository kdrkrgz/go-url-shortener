package handler

import "github.com/gofiber/fiber/v2"

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
