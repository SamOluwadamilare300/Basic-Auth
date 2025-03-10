package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Initialize database connection
	client := initializeDB()

	// Create Fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "Library API",
	})

	// Define Auth routes (public)
	AuthHandlers(app.Group("/auth"), client)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
