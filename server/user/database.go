package user

import (
	"context"
	"errors"
	"sync"

	"github.com/RajaeDev6/AlphaChat/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db = database.GetDB()
)

type DbHandler struct {
	db    *mongo.Database
	mutex sync.Mutex
}

// GetUserById retrives the user info from the databse using Id provided
func (h *DbHandler) GetUserById(Id string) (*User, error) {
	collection := db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var user User

	filter := bson.D{{"ID", Id}}

	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrives the user from the database based on the username provided
func (h *DbHandler) GetUserByUsername(Username string) (*User, error) {

	collection := db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var user User

	filter := bson.D{{"Username", Username}}
	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrives the user from the database based on the Email provided
func (h *DbHandler) GetUserByEmail(Email string) (*User, error) {

	collection := db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var user User

	filter := bson.D{{"Email", Email}}
	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}
