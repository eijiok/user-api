package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.
		PathPrefix("/api").
		Path("/users").
		Methods(http.MethodGet).
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("testing"))
		})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error on router initialization: %s", err.Error())
	}
}
