package routes

import (
	"context"
	"crud-mongodb/cmd/controller"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
	

func IntlizeRoutes(){
	r := httprouter.New()
    uc := controller.NewUserController(getClient())

    r.GET("/user/:id", uc.GetUser)
    r.POST("/user", uc.CreateUser)
    r.DELETE("/user/:id", uc.DeleteUser)
    http.ListenAndServe("localhost:6000", r)
    fmt.Print("listening on port 6000")
}

func getClient() *mongo.Client {
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
