package user

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ConfigUserPaths(router *mux.Route) {
	router.
		Path("/users").
		Methods(http.MethodGet).
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("testing"))
		})
}
