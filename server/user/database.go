package user

import (
	"context"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// DbHandler represents a handler for interacting with the MongoDB database.
type DbHandler struct {
	db    *mongo.Database
	mutex sync.Mutex
}

// NewDbHandler creates a new DbHandler instance with the provided MongoDB database connection.
func NewDbHandler(db *mongo.Database) *DbHandler {
	return &DbHandler{db: db}
}

// CreateUser function adds a user to the database
// It checks to see if user already exist
// if user exist it returns an error
func (h *DbHandler) CreateUser(user User) error {

	collection := h.db.Collection("user")

	existingUser, err := h.GetUserByUsername(user.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if existingUser != nil {
		return errors.New("User Already Taken")
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Give the new user an ObjectId
	user.ID = primitive.NewObjectID().Hex()
	user.Password = string(hashedPassword)

	// insert the new user into the database
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// GetUserById retrives the user info from the databse using Id provided
func (h *DbHandler) GetUserById(Id string) (*User, error) {

	collection := h.db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var user User

	filter := bson.D{{"_id", Id}}

	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrives the user from the database based on the username provided

func (h *DbHandler) GetUserByUsername(Username string) (*User, error) {
	var user User

	collection := h.db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	filter := bson.D{{"username", Username}}
	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrives the user from the database based on the Email provided
func (h *DbHandler) GetUserByEmail(Email string) (*User, error) {

	collection := h.db.Collection("user")
	h.mutex.Lock()
	defer h.mutex.Unlock()

	var user User

	filter := bson.D{{"email", Email}}
	// querying the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
