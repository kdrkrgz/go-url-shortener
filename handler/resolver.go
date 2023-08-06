package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kdrkrgz/go-url-shortener/repository"
	"github.com/kdrkrgz/go-url-shortener/service"
)

func ResolverHandler(repo *repository.AppRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortUrl := c.Params("short_url")
		service := &service.UrlService{
			DbRepository:    repo.DbRepository,
			CacheRepository: repo.CacheRepository,
		}

		targetUrl, err := service.FindTargetUrl(shortUrl)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		return c.Redirect(*targetUrl, fiber.StatusMovedPermanently)
	}
}
