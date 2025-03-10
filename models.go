package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	Read    Status = "read"
	Reading Status = "reading"
	ToRead  Status = "to_read"
)

// User model for MongoDB
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Password  string             `json:"password" bson:"password"`
}

// Book model for MongoDB
type Book struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"` // MongoDB ID
	Title    string             `json:"title" bson:"title"`
	Status   Status             `json:"status" bson:"status"` // Default status can be handled programmatically
	Author   string             `json:"author" bson:"author"`
	Year     int                `json:"year" bson:"year"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"` // Reference to User ID
}
