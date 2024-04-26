package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func ConnectDBandGetCLient() *mongo.Client {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
