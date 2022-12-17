package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Age      int                `bson:"age,omitempty" json:"age"`
	Email    string             `bson:"email,omitempty" json:"email"`
	Password string             `bson:"password,omitempty" json:"password"`
	Address  string             `bson:"address,omitempty" json:"address"`
}
