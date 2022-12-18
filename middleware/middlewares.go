package middleware

import (
	"github.com/eijiok/user-api/errors"
	"github.com/eijiok/user-api/interfaces"
	"net/http"
)

func CorsMiddleware(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return handler(writer, request)
	}
}

func AuthMiddleware(handler interfaces.RequestResponseErrorFunc) interfaces.RequestResponseErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			return errors.NewHttpError(http.StatusBadRequest, "Empty Authorization Header", nil)
		}
		// TODO: do a JWT token validation
		return handler(writer, request)
	}
}
