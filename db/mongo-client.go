package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoConfig struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func LoadMongoClient(ctx context.Context, mongodbURI string, mongoDatabaseName string) (*MongoConfig, error) {
	clientOptions := options.Client().ApplyURI(mongodbURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Could not get mongodb new client. Error: %s", err.Error())
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	err = client.Connect(ctxWithTimeout)
	if err != nil {
		log.Fatalf("Could not connect to mongodb database. Error: %s", err.Error())
	}

	go func() {
		<-ctx.Done()
		err = client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Error on disconecting to mongodb database. Error: %s", err.Error())
		}
	}()

	log.Printf("Connected to mongodb database!")

	clientDatabase := client.Database(mongoDatabaseName)
	return &MongoConfig{
		Client:   client,
		Database: clientDatabase,
	}, nil

}
