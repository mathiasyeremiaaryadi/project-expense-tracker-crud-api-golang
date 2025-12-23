package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRoute(app *fiber.App) {
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	app.Post("/register", RegisterHandler)
	app.Post("/login", LoginHandler)
	app.Post("/refresh-token", RefreshTokenHandler)

	app.Post("/expenses", JWTMiddleware, CreateHandler)
	app.Put("/expenses/:id", JWTMiddleware, UpdateHandler)
	app.Delete("/expenses/:id", JWTMiddleware, DeleteHandler)
	app.Get("/expenses", JWTMiddleware, GetHandler)
}
