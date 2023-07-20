package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kdrkrgz/go-url-shortener/repository"
	"github.com/kdrkrgz/go-url-shortener/resolver"
	"github.com/kdrkrgz/go-url-shortener/shortener"
)

// Shortener godoc
//
//	@Summary	Shortener
//	@Tags		Shortener
//	@Produce	json
//
// @Success	200		{object}	string
// @Router		/shortener/ [POST]
func Shortener(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Move to Service
		// target url
		var payload *shortener.Request
		var shorted resolver.Shorten
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		// check if target url is already shorted
		shorted, err := repo.FindShortedUrl(payload.TargetUrl)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		if shorted.TargetUrl == payload.TargetUrl {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"shorted_url": shorted.ShortUrl,
			})
		}
		// exp_duration := conf.Get("ExpirationDate")
		// create shorten object

		shorted = resolver.Shorten{
			TargetUrl:      payload.TargetUrl,
			ShortUrl:       shortener.GenerateShortUrl(payload.TargetUrl),
			ExpirationDate: time.Now().Add(time.Duration(time.Hour * 24)),
			CreatedAt:      time.Now(),
		}
		// insert to db
		_, err = repo.InsertShortedUrl(shorted)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Something went wrong!",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"shorted_url": shorted.ShortUrl,
		})
	}
}
