package domain

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is the database reference for all modules in domain
var db *mongo.Database
var boardCollection *mongo.Collection
var accountCollection *mongo.Collection
var accountVerificationCollection *mongo.Collection

func init() {
	mongoURL := os.Getenv("MONGO_URL")
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("agile-tools")
	initCollections()
	fmt.Println("Connected to MongoDB!")
}

func initCollections() {
	boardCollection = db.Collection("board")
	accountCollection = db.Collection("account")
	accountVerificationCollection = db.Collection("account_verification")
}
