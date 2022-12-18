package server

import (
	"context"
	"github.com/eijiok/user-api/db"
	"github.com/eijiok/user-api/infra"
	"github.com/eijiok/user-api/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitServer() error {
	config := infra.GetConfig()
	appContext := context.Background()

	mongoConfig, err := db.LoadMongoClient(appContext, config.MongodbURI, config.MongoDatabaseName)
	if err != nil {
		log.Fatalf("Database connection error. Error: %s", err.Error())
	}

	muxRouter := mux.NewRouter()

	configRoutes(mongoConfig, config, muxRouter)

	return http.ListenAndServe(":"+config.Port, muxRouter)
}

func configRoutes(mongoConfig *db.MongoConfig, config *infra.Config, muxRouter *mux.Router) {
	factory := user.GetFactory(mongoConfig)
	router := factory.GetRouter()
	pathApiPrefix := "/" + config.ApiPrefix
	pathUserApi := "/users"
	router.ConfigRoutes(muxRouter, pathApiPrefix, pathUserApi)
}
