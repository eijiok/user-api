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
		return errors.NewInternalServerErrorWithMessage(err, "Could not fetch users")
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, userList)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(err, "Error to serialize users to Json")
	}

	return nil
}

func (controller *controllerImpl) Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	user := model.User{}
	readErr := utils.ReadInJsonToStruct(request.Body, &user)
	if readErr != nil {
		return errors.NewHttpError(readErr, http.StatusBadRequest, "Error to deserialize the request body")
	}

	savedUser, err := controller.service.Save(request.Context(), &user)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(err, "Error to save the user")
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	err = utils.WriteToJson(writer, savedUser)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(err, "Saved but failed to serialize the response")
	}
	return nil
}

func (controller *controllerImpl) Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}

	user, err := controller.service.GetById(request.Context(), &objectId)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(err, "Error to fetch the document", id)
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, user)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(err, "Error to serialize the response", id)
	}
	return nil
}

func (controller *controllerImpl) Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	user := &model.User{}
	err := utils.ReadInJsonToStruct(request.Body, user)
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}

	err = controller.service.Update(request.Context(), &objectId, user)
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	writer.WriteHeader(http.StatusOK)
	return nil
}

func (controller *controllerImpl) Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}
	err := controller.service.Delete(request.Context(), &objectId)
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	writer.WriteHeader(http.StatusOK)
	return nil
}
