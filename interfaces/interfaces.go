package interfaces

import (
	"context"
	"github.com/eijiok/user-api/dto"
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RequestResponseFunc func(writer http.ResponseWriter, request *http.Request)

type RequestResponseErrorFunc func(writer http.ResponseWriter, request *http.Request) *errors.HttpError

type UserFactory interface {
	GetController() UserController
	GetService() UserService
	GetRepository() UserRepository
}

type UserController interface {
	List(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError
	Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError
}

type UserService interface {
	List(ctx context.Context) ([]dto.User, error)
	GetById(ctx context.Context, id *primitive.ObjectID) (*dto.User, error)
	Save(ctx context.Context, user *model.User) (*dto.User, error)
	Update(ctx context.Context, id *primitive.ObjectID, user *model.User) error
	Delete(ctx context.Context, id *primitive.ObjectID) error
}

type UserRepository interface {
	List(context.Context) ([]model.User, error)
	GetById(ctx context.Context, id *primitive.ObjectID) (*model.User, error)
	Save(ctx context.Context, user *model.User) (*primitive.ObjectID, error)
	Update(ctx context.Context, user *model.User) (int64, error)
	Delete(ctx context.Context, id *primitive.ObjectID) (int64, error)
}
