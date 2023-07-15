package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Initialize initializes the database connection.
// This function establishes a connection to MongoDB using the MongoDB driver.
func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	DB_NAME := os.Getenv("DB_NAME")

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Access the MongoDB database
	db = client.Database(DB_NAME)
}

// GetDB returns the MongoDB database connection.
// This function allows other functions or models to access and interact with the MongoDB database.
func GetDB() *mongo.Database {
	return db
}
