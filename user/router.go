package user

import (
	"github.com/eijiok/user-api/interfaces"
	"github.com/eijiok/user-api/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type userRouterImpl struct {
	controller          interfaces.UserController
	secureMiddlewares   []interfaces.MiddlewareFunc
	insecureMiddlewares []interfaces.MiddlewareFunc
}

func (routes *userRouterImpl) ConfigRoutes(apiRouter *mux.Router, pathPrefix string, pathUserApi string) {
	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodGet).
		HandlerFunc(routes.setSecureHandler(routes.controller.List))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi).
		Methods(http.MethodPost).
		HandlerFunc(routes.setInsecureHandler(routes.controller.Create))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodGet).
		HandlerFunc(routes.setSecureHandler(routes.controller.Read))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodPut).
		HandlerFunc(routes.setSecureHandler(routes.controller.Update))

	apiRouter.
		PathPrefix(pathPrefix).
		Path(pathUserApi + "/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(routes.setSecureHandler(routes.controller.Delete))
}

func (routes *userRouterImpl) setSecureHandler(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseFunc {
	return utils.RequestResponseErrorHandler(routes.addMiddlewares(handler, routes.secureMiddlewares))
}

func (routes *userRouterImpl) setInsecureHandler(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseFunc {
	return utils.RequestResponseErrorHandler(routes.addMiddlewares(handler, routes.insecureMiddlewares))
}

func (routes *userRouterImpl) addMiddlewares(handler interfaces.RequestResponseErrorFunc, middlewares []interfaces.MiddlewareFunc) interfaces.RequestResponseErrorFunc {
	currHandler := handler
	for _, middleware := range middlewares {
		currHandler = middleware(currHandler)
	}
	return currHandler
}

func NewUserRouter(
	controller interfaces.UserController,
	secureMiddlewares []interfaces.MiddlewareFunc,
	insecureMiddlewares []interfaces.MiddlewareFunc,
) interfaces.UserRouter {
	return &userRouterImpl{
		controller:          controller,
		secureMiddlewares:   secureMiddlewares,
		insecureMiddlewares: insecureMiddlewares,
	}
}
