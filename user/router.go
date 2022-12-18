package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type userRouterImpl struct {
	controller  interfaces.UserController
	middlewares []interfaces.MiddlewareFunc
}

func (routes *userRouterImpl) ConfigRoutes(apiRouter *mux.Router, pathPrefix string, pathUserApi string) {
	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodGet).
		HandlerFunc(routes.setHandler(routes.controller.List))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodPost).
		HandlerFunc(routes.setHandler(routes.controller.Create))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodGet).
		HandlerFunc(routes.setHandler(routes.controller.Read))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodPut).
		HandlerFunc(routes.setHandler(routes.controller.Update))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(routes.setHandler(routes.controller.Delete))
}

func (routes *userRouterImpl) setHandler(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseFunc {

	return utils.RequestResponseErrorHandler(routes.addMiddlewares(handler))
}

func (routes *userRouterImpl) addMiddlewares(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseErrorFunc {
	currHandler := handler
	for _, middleware := range routes.middlewares {
		currHandler = middleware(currHandler)
	}
	return currHandler
}

func NewUserRouter(controller interfaces.UserController, middlewares ...interfaces.MiddlewareFunc) interfaces.UserRouter {
	return &userRouterImpl{
		controller:  controller,
		middlewares: middlewares,
	}
}
