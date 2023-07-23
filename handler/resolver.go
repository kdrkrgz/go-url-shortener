package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kdrkrgz/go-url-shortener/repository"
)

func ResolverHandler(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortUrl := c.Params("short_url")
		targetUrl, err := repo.FindUrl("short_url", shortUrl)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		return c.Redirect(*targetUrl, fiber.StatusMovedPermanently)
	}
}
