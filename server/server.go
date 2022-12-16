package server

import (
	"github.com/eijiok/user-api/user"
	"github.com/gorilla/mux"
	"net/http"
)

func InitServer(pathPrefix string, port string) error {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/" + pathPrefix)
	user.ConfigUserPaths(apiRouter)
	return http.ListenAndServe(":"+port, router)

}
