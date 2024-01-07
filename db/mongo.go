// db/mongodb.go
package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// InitMongoDB initializes the MongoDB connection
func InitMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	collection = client.Database("mydb").Collection("people")
	fmt.Println("Connected to MongoDB!")

	return nil
}

// GetCollection returns the MongoDB collection instance
func GetCollection() *mongo.Collection {
	return collection
}
