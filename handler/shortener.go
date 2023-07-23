package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kdrkrgz/go-url-shortener/conf"
	"github.com/kdrkrgz/go-url-shortener/repository"
	"github.com/kdrkrgz/go-url-shortener/resolver"
	"github.com/kdrkrgz/go-url-shortener/shortener"
)

// Shortener godoc
//
//	@Router		/shortener/ [POST]
func ShortenerHandler(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// target url
		var payload *shortener.Request
		var shorted resolver.ShortUrl
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		// check if target url is already shorted
		shortedUrl, _ := repo.FindUrl("target_url", payload.TargetUrl)
		if shortedUrl != nil {
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"shorted_url": fmt.Sprintf("%s/%s", conf.Get("App.Domain"), *shortedUrl),
			})
		}
		// create shorten object
		shorted = resolver.ShortUrl{
			TargetUrl:      payload.TargetUrl,
			ShortUrl:       shortener.GenerateShortUrl(),
			ExpirationDate: time.Now().Add(time.Duration(time.Hour * 24)),
			CreatedAt:      time.Now(),
		}
		// insert to db
		_, errInsert := repo.InsertShortedUrl(shorted)
		if errInsert != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Something went wrong!",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"shorted_url": fmt.Sprintf("%s/%s", conf.Get("App.Domain"), shorted.ShortUrl),
		})
	}
}
