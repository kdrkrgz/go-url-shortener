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
	tasks "github.com/kdrkrgz/go-url-shortener/pkg/tasks"
	"github.com/kdrkrgz/go-url-shortener/repository"
)

type Application struct {
	app  *fiber.App
	repo *repository.Repository
}

func (a *Application) Register() {
	a.app.Get("/", handler.RedirectSwagger)
	a.app.Get("/healthcheck", handler.HealthCheck)
	a.app.Post("/shorten", handler.ShortenerHandler(a.repo))
	a.app.Get("/:short_url", handler.ResolverHandler(a.repo))
	route := a.app.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

// @title						GoUrlShortener API
// @version					    1.0
// @description				    Swagger for GoUrlShortener app
// @host						localhost:8000
// @BasePath					/
// @schemes					    http
// @license.name				Apache License, Version 2.0 (the "License")
func main() {
	repo := repository.New()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
		AllowCredentials: true,
	}))
	// app.Use(limiter.New(limiter.Config{Max: 2, Expiration: 1 * time.Minute}))
	application := &Application{app: app, repo: repo}
	application.Register()
	tasks.RunTasks()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		log.Logger().Info("application gracefully shutting down..")
		_ = app.Shutdown()
	}()

	if err := app.Listen(fmt.Sprintf(":%v", conf.Get("App.AppPort"))); err != nil {
		log.Logger().Panic(fmt.Sprintf("App Err: %s", err))
	}
}
