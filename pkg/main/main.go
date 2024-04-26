// package main

// import (
// 	"crud-mongodb/cmd/controller"
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// )

// func main() {
// 	r := httprouter.New()
// 	uc := controller.NewUserController(getSession())

// 	r.GET("/user/:id", uc.GetUser)
// 	r.POST("/user", uc.CreateUser)
// 	r.DELETE("/user/:id", uc.DeleteUser)
// 	http.ListenAndServe("localhost:6000", r)
// 	fmt.Print("listing on port ")

// }

// func getSession() *mgo.Session {
// 	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
// 	if err != nil {
// 		fmt.Println("Failed to connect to MongoDB:", err)
// 		panic(err)
// 	}
// 	return session
// }

package main

import (
	"crud-mongodb/cmd/routes"
)

func main() {
	routes.IntlizeRoutes()
	// r := httprouter.New()
	// uc := controller.NewUserController(getClient())

	// r.GET("/user/:id", uc.GetUser)
	// r.POST("/user", uc.CreateUser)
	// r.DELETE("/user/:id", uc.DeleteUser)
	// http.ListenAndServe("localhost:6000", r)
	// fmt.Print("listening on port 6000")
}

// func getClient() *mongo.Client {
//     // Set client options
//     clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

//     // Connect to MongoDB
//     client, err := mongo.Connect(context.Background(), clientOptions)
//     if err != nil {
//         fmt.Println("Failed to connect to MongoDB:", err)
//         panic(err)
//     }

//     // Ping the MongoDB server to check if the connection was successful
//     err = client.Ping(context.Background(), nil)
//     if err != nil {
//         fmt.Println("Failed to ping MongoDB:", err)
//         panic(err)
//     }

//     fmt.Println("Connected to MongoDB!")
//     return client
// }
