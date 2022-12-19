package interfaces

import (
	"context"
	"github.com/eijiok/user-api/dto"
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RequestResponseFunc func(writer http.ResponseWriter, request *http.Request)

type RequestResponseErrorFunc func(writer http.ResponseWriter, request *http.Request) *errors.HttpError

type MiddlewareFunc func(handler RequestResponseErrorFunc) RequestResponseErrorFunc

type UserFactory interface {
	GetController() UserController
	GetService() UserService
	GetRepository() UserRepository
	GetRouter() UserRouter
}

type UserController interface {
	List(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError
}

type UserService interface {
	List(ctx context.Context) ([]dto.UserResponse, error)
	GetById(ctx context.Context, id *primitive.ObjectID) (*dto.UserResponse, error)
	Save(ctx context.Context, user *dto.UserRequest) (*dto.UserResponse, error)
	Update(ctx context.Context, id *primitive.ObjectID, user *dto.UserRequest) error
	Delete(ctx context.Context, id *primitive.ObjectID) error
}

type UserRepository interface {
	List(context.Context) ([]model.User, error)
	GetById(ctx context.Context, id *primitive.ObjectID) (*model.User, error)
	Save(ctx context.Context, user *model.User) (*primitive.ObjectID, error)
	Update(ctx context.Context, user *model.User) (int64, error)
	Delete(ctx context.Context, id *primitive.ObjectID) (int64, error)
}

type UserRouter interface {
	ConfigRoutes(apiRouter *mux.Router, pathPrefix string, pathUserApi string)
}
