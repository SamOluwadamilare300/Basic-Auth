package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthHandlers defines the authentication routes
func AuthHandlers(route fiber.Router, client *mongo.Client) {
	// Get the user collection from MongoDB
	collection := client.Database("library").Collection("users")

	// Register route
	route.Post("/register", func(c *fiber.Ctx) error {
		// Parse incoming user data from JSON
		var user struct {
			Firstname string `json:"firstname"`
			Email     string `json:"email"`
			Lastname  string `json:"lastname"`
			Password  string `json:"password"`
		}

		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString("Invalid input")
		}

		// Insert new user into MongoDB
		_, err := collection.InsertOne(c.Context(), bson.D{
			{Key: "firstname", Value: user.Firstname},
			{Key: "lastname", Value: user.Lastname},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
		})
		if err != nil {
			log.Println("Error inserting user:", err)
			return c.Status(500).SendString("Failed to register")
		}

		return c.Status(201).SendString("User registered successfully")
	})

	// Login route
	route.Post("/login", func(c *fiber.Ctx) error {
		// Parse login credentials
		var user struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString("Invalid input")
		}

		// Find user in MongoDB
		var result struct {
			Email    string `bson:"email"`
			Password string `bson:"password"`
		}

		err := collection.FindOne(c.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(404).SendString("User not found")
			}
			log.Println("Error fetching user:", err)
			return c.Status(500).SendString("Failed to login")
		}

		// Check password (for simplicity, not using hashing here)
		if result.Password != user.Password {
			return c.Status(401).SendString("Invalid password")
		}

		return c.SendString("Login successful")
	})
}
