package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	conf "github.com/kdrkrgz/go-url-shortener/conf"
	handler "github.com/kdrkrgz/go-url-shortener/handler"
	log "github.com/kdrkrgz/go-url-shortener/pkg/logger"
	"github.com/kdrkrgz/go-url-shortener/repository"
)

type Application struct {
	app        *fiber.App
	collection *repository.Repository
}

func (a *Application) Register() {
	a.app.Get("/", handler.RedirectSwagger)
	a.app.Get("/healthcheck", handler.HealthCheck)
	a.app.Post("/shorten", handler.Shortener(a.collection))
	// a.app.Get("/resolver", handler.Redirect(a.collection))
	route := a.app.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

func main() {
	collection := repository.New()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
		AllowCredentials: true,
	}))
	application := &Application{app: app, collection: collection}
	application.Register()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		log.Logger().Info("application gracefully shutting down..")
		_ = app.Shutdown()
	}()

	if err := app.Listen(fmt.Sprintf(":%v", conf.Get("AppPort"))); err != nil {
		log.Logger().Panic(fmt.Sprintf("App Err: %s", err))
	}
}
