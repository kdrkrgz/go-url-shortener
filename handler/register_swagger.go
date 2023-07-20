package handler

import "github.com/gofiber/fiber/v2"

func RedirectSwagger(ctx *fiber.Ctx) error {
	return ctx.Redirect("/swagger/index.html")
}
