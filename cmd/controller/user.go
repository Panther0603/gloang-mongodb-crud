// package controller

// import (
// 	"fmt"
// 	"net/http"

// 	"encoding/json"

// 	"crud-mongodb/cmd/models"

// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type UserController struct {
// 	session *mgo.Session
// }

// func NewUserController(se *mgo.Session) *UserController {

// 	return &UserController{session: se}
// }

// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	id := p.ByName("id")

// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	oid := bson.ObjectIdHex(id)
// 	u := models.User{}

// 	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
// 		w.WriteHeader(404)
// 		fmt.Println(err)
// 		return
// 	}

// 	uj, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)

// 	fmt.Fprintf(w, "%s\n", uj)

// }

// func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// 	u := models.User{}

// 	json.NewDecoder(r.Body).Decode(&u)
// 	u.Id = bson.NewObjectId()
// 	uc.session.DB("mongo-golang").C("users").Insert(u)

// 	uj, err := json.Marshal(u)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintf(w, "%s\n", uj)

// }

// func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	id := p.ByName("id")
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	oid := bson.ObjectIdHex(id)

// 	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
// 		w.WriteHeader(404)
// 	}
// 	w.WriteHeader(http.StatusOK)

// 	fmt.Fprint(w, "Deleted user", oid, "\n")

// }

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"crud-mongodb/cmd/models"

	"context"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client: client}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Println(id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user ID: %v", err)
		return
	}

	var user models.User
	collection := uc.client.Database("mongo-golang").Collection("users")
	filter := bson.M{"_id": objID}

	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "User not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user: %v", err)
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	user.Id = primitive.NewObjectID()
	collection := uc.client.Database("mongo-golang").Collection("users")
	_, err := collection.InsertOne(context.Background(), &user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting user: %v", err)
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users")
	filter := bson.M{"_id": id}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s deleted successfully", id)
}

func (uc UserController) GetALlUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	collection := uc.client.Database("mongo-golang").Collection("users")

	userPas, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching users: %v", err)
		return
	}

	var users []models.User
	if err := userPas.All(context.Background(), &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while fetching users : %v", err)
	}

	responseJson, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)

}
