package user

import (
	"github.com/eijiok/user-api/dto"
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/interfaces"
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

// swagger:response ListUserResponse
type ListUserResponse struct {
	// in: body
	Body []dto.UserResponse
}

// List
/**
 * swagger:route GET /api/users Users ListUsers
 *
 * Fetch the registered users
 *
 * Produces:
 * - application/json
 *
 * Responses:
 *    200: ListUserResponse
 *    500:
 */
func (controller *controllerImpl) List(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	userList, err := controller.service.List(request.Context())
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(&err, "Could not fetch users")
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, userList)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(&err, "Error to serialize users to Json")
	}

	return nil
}

// swagger:parameters CreateUserRequest
type CreateUserRequest struct {
	// Request to create a new User
	//
	// required: true
	// in: body
	CreateUserRequest dto.UserRequest
}

// swagger:response CreateUserResponse
type CreateUserResponse struct {
	// in: body
	Body dto.UserResponse
}

// Create
/**
 * swagger:route POST /api/users Users CreateUserRequest
 *
 * Creates a user from the request parameters
 *
 * Produces:
 * - application/json
 *
 * Responses:
 *    201: CreateUserResponse
 *    400:
 *    401:
 *    500:
 */
func (controller *controllerImpl) Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	user := dto.UserRequest{}
	readErr := utils.ReadInJsonToStruct(request.Body, &user)
	if readErr != nil {
		return errors.NewHttpError(http.StatusBadRequest, "Error to deserialize the request body", &readErr)
	}

	savedUser, err := controller.service.Save(request.Context(), &user)
	if err != nil {
		errorStatusCode := http.StatusInternalServerError
		if _, ok := err.(*errors.ValidationError); ok {
			errorStatusCode = http.StatusBadRequest
		}
		return errors.NewHttpError(errorStatusCode, "Error to save the user", &err)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	err = utils.WriteToJson(writer, savedUser)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(&err, "Saved but failed to serialize the response")
	}
	return nil
}

// swagger:response ReadUserResponse
type ReadUserResponse struct {
	// in: body
	Body dto.UserResponse
}

// Read
/**
 * swagger:route GET /api/users/{userId} Users ReadUsers
 *
 * Fetches a user from the path parameter
 *
 * Parameters:
 *   + name: userId
 *     in: path
 *     description: the user ID
 *     required: true
 *     type: string
 *
 * Produces:
 * - application/json
 *
 * Responses:
 *    200: ReadUserResponse
 *    400:
 *    500:
 */
func (controller *controllerImpl) Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}

	user, err := controller.service.GetById(request.Context(), &objectId)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(&err, "Error to fetch the document", id)
	}

	writer.WriteHeader(http.StatusOK)
	err = utils.WriteToJson(writer, user)
	if err != nil {
		return errors.NewInternalServerErrorWithMessage(&err, "Error to serialize the response", id)
	}
	return nil
}

// swagger:parameters UpdateUserRequest
type UpdateUserRequest struct {
	// Request to update a new User
	//
	// required: true
	// in: body
	UpdateUserRequest dto.UserRequest
}

// swagger:response UpdateUserResponse
type UpdateUserResponse struct {
	// in: body
	Body dto.UserResponse
}

// Update
/**
 * swagger:route PUT /api/users/{userId} Users UpdateUserRequest
 *
 * Updates a user from the request parameters
 *
 * Parameters:
 *   + name: userId
 *     in: path
 *     description: the user ID
 *     required: true
 *     type: string
 *
 * Produces:
 * - application/json
 *
 * Responses:
 *    200: UpdateUserResponse
 *    400:
 *    401:
 *    500:
 */
func (controller *controllerImpl) Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	user := &dto.UserRequest{}
	err := utils.ReadInJsonToStruct(request.Body, user)
	if err != nil {
		return errors.NewInternalServerError(&err)
	}

	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}

	err = controller.service.Update(request.Context(), &objectId, user)
	if err != nil {
		errorStatusCode := http.StatusInternalServerError
		if _, ok := err.(*errors.ValidationError); ok {
			errorStatusCode = http.StatusBadRequest
		}
		return errors.NewHttpError(errorStatusCode, "Error to update the user", &err)
	}

	writer.WriteHeader(http.StatusOK)
	return nil
}

// Delete
/**
 * swagger:route DELETE /api/users/{userId} Users ReadUsers
 *
 * Deletes a user from the path parameters
 *
 * Parameters:
 *   + name: userId
 *     in: path
 *     description: the user ID
 *     required: true
 *     type: string
 *
 * Produces:
 * - application/json
 *
 * Responses:
 *    200:
 *    400:
 *    500:
 */
func (controller *controllerImpl) Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	id := utils.ReadParam(request, "id")
	objectId, httpError := utils.IdToObjectId(id)
	if httpError != nil {
		return httpError
	}
	err := controller.service.Delete(request.Context(), &objectId)
	if err != nil {
		return errors.NewInternalServerError(&err)
	}
	writer.WriteHeader(http.StatusOK)
	return nil
}
