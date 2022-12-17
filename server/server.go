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

	router := mux.NewRouter()

	user.InitConf(config.ApiPrefix, router, user.GetFactory(mongoConfig))

	return http.ListenAndServe(":"+config.Port, router)
}
