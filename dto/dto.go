package dto

import (
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Age     int                `json:"age"`
	Email   string             `json:"email"`
	Address string             `json:"address"`
}

func (dto *User) ToUserModel() *model.User {
	return &model.User{
		ID:      dto.ID,
		Name:    dto.Name,
		Age:     dto.Age,
		Email:   dto.Email,
		Address: dto.Address,
	}
}

func (dto *User) FromUserModel(userModel *model.User) {
	dto.ID = userModel.ID
	dto.Name = userModel.Name
	dto.Age = userModel.Age
	dto.Email = userModel.Email
	dto.Address = userModel.Address
}
