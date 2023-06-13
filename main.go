package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// middleware
	app.Use(logger.New())

	// routes
	registerRoutes(app)

	// start
	app.Listen(":8080")
}
