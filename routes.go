package main

import "github.com/gofiber/fiber/v2"

// CRUD: create read update delete
// HTTP: GEt POST DELETE PUT PATCH
func registerRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	{
		userGroup.Post("/", createUser)          // done
		userGroup.Get("/", getAllUsers)          // done
		userGroup.Get("/:userId", getUser)       // done
		userGroup.Put("/:userId", update)        // complete update
		userGroup.Patch("/:userId", toggleAdmin) // partial update
		userGroup.Delete("/:userId", delete)     // done
	}
}
