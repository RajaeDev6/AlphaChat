package database

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db      *mongo.Database
	dbMutex sync.Mutex
)

// Initialize initializes the database connection.
// This function establishes a connection to MongoDB using the MongoDB driver.
func Initialize() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	DB_NAME := os.Getenv("DB_NAME")

	// Create a new context with a timeout (e.g., 30 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {

		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Access the MongoDB database
	dbMutex.Lock()
	db = client.Database(DB_NAME)
	dbMutex.Unlock()

	return nil
}

// GetDB returns the MongoDB database connection.
// This function allows other functions or models to access and interact with the MongoDB database.
func GetDB() *mongo.Database {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return db
}
