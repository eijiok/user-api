package model

type User struct {
	ID       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Age      int    `bson:"age" json:"age"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Address  string `bson:"address" json:"address"`
}
