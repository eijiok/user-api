package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var config *Config

type Config struct {
	Port              string
	ApiPrefix         string
	MongodbURI        string
	MongoDatabaseName string
}

func GetConfig() *Config {
	if config == nil {
		initConfig()
	}

	return config
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		config = &Config{
			Port:              "8080",
			ApiPrefix:         "api",
			MongodbURI:        "mongodb://root:root@localhost:27017/admin",
			MongoDatabaseName: "api-user-mongo",
		}
	} else {
		config = &Config{
			Port:              os.Getenv("PORT"),
			ApiPrefix:         os.Getenv("API_PREFIX"),
			MongodbURI:        os.Getenv("MONGODB_URI"),
			MongoDatabaseName: os.Getenv("DATABASE_NAME"),
		}
	}
}
