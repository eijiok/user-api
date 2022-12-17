package server

import (
	"github.com/eijiok/user-api/user"
	"github.com/gorilla/mux"
	"net/http"
)

func InitServer(pathPrefix string, port string) error {
	router := mux.NewRouter()
	user.InitConf(pathPrefix, router)
	return http.ListenAndServe(":"+port, router)
}
