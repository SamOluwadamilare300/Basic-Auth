package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware that returns a message for non-authenticated routes
func NoAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Welcome to the Library API!",
		})
	}
}

func main() {
	// Create Fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "Library API",
	})

	// Root route to display message when calling `http://127.0.0.1:3000`
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Welcome to the Library API. Use the /auth routes for authentication.",
		})
	})

	// Initialize database connection
	client := initializeDB()

	// Define Auth routes (public)
	AuthHandlers(app.Group("/auth"), client)

	// Apply middleware globally (optional)
	app.Use(NoAuthMiddleware())

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
