package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//swagger:model User
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Birthday  time.Time          `bson:"birthday" json:"birthday"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Address   string             `bson:"address" json:"address"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}
