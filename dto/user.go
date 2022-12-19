package dto

import (
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// swagger:parameters CreateUserRequest
type CreateUserRequest struct {
	// Request to create a new User
	//
	// required: true
	// in: body
	CreateUserRequest UserRequest
}

type UserRequest struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Address  string    `json:"address"`
}

func (req *UserRequest) ToUser() model.User {
	return model.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Birthday: req.Birthday,
		Address:  req.Address,
	}
}

// swagger:response CreateUserResponse
type CreateUserResponse struct {
	// in: body
	Body UserResponse
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Address   string             `json:"address"`
	Birthday  time.Time          `json:"birthday"`
	Email     string             `json:"email"`
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
