package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDBandGetCLient() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Print(os.Getenv("DB_URI"))
	// clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}

	// Ping the MongoDB server to check if the connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func CreateWithUnique(collectionName string) {
	client := ConnectDBandGetCLient()
	collection := client.Database("mongo-golang").Collection(collectionName)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		fmt.Print("Got some error while creating index uniquely:", err)
	}
}
