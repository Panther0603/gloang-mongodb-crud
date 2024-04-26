package routes

import (
	"crud-mongodb/cmd/controller"
	"crud-mongodb/cmd/database"
	"crud-mongodb/cmd/models"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

func IntlizeRoutes() {
	r := httprouter.New()
	uc := controller.NewUserController(getClient())
	models.CreateWithUnique()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/user", uc.GetALlUsers)
	http.ListenAndServe("localhost:6000", r)
	fmt.Print("listening on port 6000")
}

func getClient() *mongo.Client {
	return database.ConnectDBandGetCLient()
}
