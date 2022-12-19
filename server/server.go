package server

import (
	"context"
	"github.com/eijiok/user-api/db"
	"github.com/eijiok/user-api/infra"
	"github.com/eijiok/user-api/user"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func InitServer() error {
	config := infra.GetConfig()
	appContext := context.Background()

	mongoConfig, err := db.LoadMongoClient(appContext, config.MongodbURI, config.MongoDatabaseName)
	if err != nil {
		log.Fatalf("Database connection error. Error: %s", err.Error())
	}

	muxRouter := mux.NewRouter()

	configSwaggerRoutes(muxRouter)

	configAppRoutes(mongoConfig, config, muxRouter)
	log.Printf("The server started on port %s ", config.Port)

	return http.ListenAndServe(":"+config.Port, muxRouter)
}

func configSwaggerRoutes(muxRouter *mux.Router) {
	swaggerOpts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	swaggerMiddleware := middleware.SwaggerUI(swaggerOpts, nil)
	muxRouter.Handle("/docs", swaggerMiddleware)
	muxRouter.
		Path("/swagger.json").
		Methods(http.MethodGet).
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fileBytes, _ := os.ReadFile("swagger.json")
			writer.WriteHeader(http.StatusOK)
			writer.Header().Set("Content-Type", "application/octet-stream")
			_, _ = writer.Write(fileBytes)
		})
}

func configAppRoutes(mongoConfig *db.MongoConfig, config *infra.Config, muxRouter *mux.Router) {
	factory := user.GetFactory(mongoConfig)
	router := factory.GetRouter()
	pathApiPrefix := "/" + config.ApiPrefix
	pathUserApi := "/users"
	router.ConfigRoutes(muxRouter, pathApiPrefix, pathUserApi)
}
