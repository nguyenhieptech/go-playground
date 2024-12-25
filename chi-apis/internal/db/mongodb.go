package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.Background()

func InitMongoDB() (*mongo.Client, *mongo.Collection) {
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err := client.Ping(Ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("Connected to MongoDB")
	collection := client.Database("golang_db").Collection("todos")
	return client, collection
}
