package dto

import (
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Address   string             `json:"address"`
	Birthday  time.Time          `json:"birthday"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func (dto *UserResponse) FromUserModel(userModel *model.User) {
	dto.ID = userModel.ID
	dto.Name = userModel.Name
	dto.Address = userModel.Address
	dto.Birthday = userModel.Birthday
	dto.Email = userModel.Email
	dto.CreatedAt = userModel.CreatedAt
	dto.UpdatedAt = userModel.UpdatedAt
}
