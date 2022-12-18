package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type userRouterImpl struct {
	controller interfaces.UserController
}

func (routes *userRouterImpl) ConfigRoutes(apiRouter *mux.Router, pathPrefix string, pathUserApi string) {
	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodGet).
		HandlerFunc(utils.RequestResponseErrorHandler(routes.controller.List))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodPost).
		HandlerFunc(utils.RequestResponseErrorHandler(routes.controller.Create))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodGet).
		HandlerFunc(utils.RequestResponseErrorHandler(routes.controller.Read))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodPut).
		HandlerFunc(utils.RequestResponseErrorHandler(routes.controller.Update))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(utils.RequestResponseErrorHandler(routes.controller.Delete))
}

func NewUserRouter(controller interfaces.UserController) interfaces.UserRouter {
	return &userRouterImpl{
		controller: controller,
	}
}
