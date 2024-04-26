package models

import (
	"time"

	"crud-mongodb/cmd/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Gender      string             `json:"gender" bson:"gender"`
	Age         int                `json:"age" bson:"age"`
	CreatedDate time.Time          `json:"createdDate" bson:"creatddate"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"creatdDate"`
	Email       string             `json:"email" bson:"email"`
}

func CreateWithUnique() {
	database.CreateWithUnique("users")
}
