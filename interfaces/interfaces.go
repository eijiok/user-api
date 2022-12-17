package interfaces

import (
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/model"
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
	List() ([]model.User, error)
	GetById(id string) (model.User, error)
	Save(user model.User) (model.User, error)
	Update(user model.User) error
	Delete(id string) error
}

type UserRepository interface {
	List() ([]model.User, error)
	GetById(id string) (model.User, error)
	Save(user model.User) (model.User, error)
	Update(user model.User) error
	Delete(id string) error
}
