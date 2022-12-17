package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func InitConf(pathPrefix string, apiRouter *mux.Router, userFactory interfaces.UserFactory) {

	controller := userFactory.GetController()
	pathPrefix = "/" + pathPrefix
	pathUserApi := "/users"
	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodGet).
		HandlerFunc(utils.RequestResponseErrorHandler(controller.List))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodPost).
		HandlerFunc(utils.RequestResponseErrorHandler(controller.Create))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodGet).
		HandlerFunc(utils.RequestResponseErrorHandler(controller.Read))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodPut).
		HandlerFunc(utils.RequestResponseErrorHandler(controller.Update))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(utils.RequestResponseErrorHandler(controller.Delete))
}
