package domain

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charlesmarvin/agile-tools/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var boardsCollection *mongo.Collection

func init() {
	mongoURL := os.Getenv("MONGO_URL")
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("agile-tools")
	boardsCollection = db.Collection("boards")
	fmt.Println("Connected to MongoDB!")
}

// SaveBoard creates a new board
func SaveBoard(board models.Board) (*models.Board, error) {
	_, err := boardsCollection.InsertOne(context.TODO(), board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}

// GetBoard gets a board by ID
func GetBoard(id string) (*models.Board, error) {
	var board models.Board
	err := boardsCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&board)
	return &board, err
}
