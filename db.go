package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initializeDB connects to MongoDB and returns a MongoDB client instance.
func initializeDB() *mongo.Client {
	
	uri := "mongodb+srv://Library74:A4OTUvs11JoSQbJp@library.y35qs.mongodb.net"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
