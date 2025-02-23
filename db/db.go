package db

import (
	"context"
	"fmt"
	"log"
	"narad/env"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// # MongoDB Client, Database and Collection Variables
var client *mongo.Client
var db *mongo.Database
var collection *mongo.Collection

// # Initialize MongoDB Connection
func init() {
	// # Load Env Variables
	env.LoadEnv()

	// # Get DB User and DB Password
	dbUser := env.GetEnv("DB_USER", "User")
	dbPass := env.GetEnv("DB_PASS", "Password")

	// # Build MongoDB URI with DB User and DB Password
	mongoURI := os.Getenv("MONGO_URI")
	mongoURI = strings.Replace(mongoURI, "{DB_USER}", dbUser, 1)
	mongoURI = strings.Replace(mongoURI, "{DB_PASS}", dbPass, 1)

	// # Get DB Name and Collection Name
	dbName := env.GetEnv("DB_NAME", "DB")
	collectionName := env.GetEnv("COLLECTION_NAME", "Collection")

	// # Error Object
	var err error

	// # Context
	ctx := context.TODO()

	// # MongoDB Client Options with MongoDB URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// # Connect to MongoDB Server with MongoDB Client Options
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		// # MongoDB Connection Error
		log.Fatal("🚫 MongoDB Connection Error: ", err)
	}

	// # Ping the MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		// # MongoDB Ping Error
		log.Fatal("🚫 MongoDB Ping Error: ", err)
	}

	// # Set the MongoDB Database and Collection
	db = client.Database(dbName)
	collection = db.Collection(collectionName)

	// # Server Name
	serverName := "Narad Service"

	// # Server Connection Successful Message
	fmt.Printf("🕸️  %s Server Connected!\n", serverName)

	// # MongoDB Connection Successful Message
	fmt.Printf("🕸️  %s Database Connected!\n", serverName)
}

// # Get MongoDB Collection
func GetCollection() *mongo.Collection {
	return collection
}

// # Get Collection By Name
func GetCollectionByName(collectionName string) *mongo.Collection {
	collection = db.Collection(collectionName)
	return collection
}
