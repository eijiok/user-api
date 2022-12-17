package user

import (
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/model"
	"github.com/eijiok/user-api/utils"
	"net/http"
)

type controllerImpl struct {
	service interfaces.UserService
}

func NewControllerImpl(service interfaces.UserService) interfaces.UserController {
	return &controllerImpl{
		service: service,
	}
}

func (controller *controllerImpl) List(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	userList, err := controller.service.List(request.Context())
	if err != nil {
		return errors.NewHttpError(err)
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, userList)
	if err != nil {
		return errors.NewHttpError(err)
	}

	return nil
}

func (controller *controllerImpl) Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	user := model.User{}
	readErr := utils.ReadInJsonToStruct(request.Body, &user)
	if readErr != nil {
		return errors.NewHttpError(readErr)
	}

	savedUser, err := controller.service.Save(request.Context(), &user)
	if err != nil {
		return errors.NewHttpError(err)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	err = utils.WriteToJson(writer, savedUser)
	if err != nil {
		return errors.NewHttpError(err)
	}
	return nil
}

func (controller *controllerImpl) Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	user, err := controller.service.GetById(request.Context(), id)
	if err != nil {
		return errors.NewHttpError(err)
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, user)
	if err != nil {
		return errors.NewHttpError(err)
	}
	return nil
}

func (controller *controllerImpl) Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	user := &model.User{}
	err := utils.ReadInJsonToStruct(request.Body, user)
	if err != nil {
		return errors.NewHttpError(err)
	}

	err = controller.service.Update(request.Context(), id, user)
	if err != nil {
		return errors.NewHttpError(err)
	}

	writer.WriteHeader(http.StatusOK)
	return nil
}

func (controller *controllerImpl) Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	err := controller.service.Delete(request.Context(), id)
	if err != nil {
		return errors.NewHttpError(err)
	}
	writer.WriteHeader(http.StatusOK)
	return nil
}
