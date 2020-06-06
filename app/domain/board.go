package domain

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var boardsCollection *mongo.Collection

// BoardType is a enumeration of board types [Fibonacci|Standard|Tshirt]
type BoardType string

// BoardID the board id
type BoardID string

// Board Model object for a board
type Board struct {
	Name     string    `bson:"name" json:"name"`
	Type     BoardType `bson:"type" json:"type"`
	ID       BoardID   `bson:"_id" json:"id"`
	Passcode string    `bson:"passcode" json:"passcode"`
}

// ContactType is a enumeration of contact types [Phone|Email]
type ContactType string

// Contact represents either a phone or email contact
type Contact struct {
	Type  ContactType `bson:"contactType" json:"contactType"`
	Value string      `bson:"value" json:"value"`
}

// CreateBoardRequest is the input for creating a new board
type CreateBoardRequest struct {
	Name     string    `json:"name"`
	Type     BoardType `json:"type"`
	Passcode string    `json:"passcode"`
	Members  []Contact `json:"members"`
}

// NewBoardID Creates a new random BoardID
func NewBoardID() BoardID {
	guid := strings.Replace(uuid.New().String(), "-", "", -1)
	guid = strings.ToUpper(guid)
	return BoardID("brd_" + guid)
}

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

// Save creates a new board
func (board *Board) Save() (*Board, error) {
	_, err := boardsCollection.InsertOne(context.TODO(), board)
	if err != nil {
		return nil, err
	}
	return board, nil
}

// Get gets a board by ID
func (board *Board) Get(id string) (*Board, error) {
	err := boardsCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(board)
	return board, err
}
