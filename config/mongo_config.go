package config

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var newClient *mongo.Client
var databaseName string
var once sync.Once

func NewClientDB() (*mongo.Client, string, error) {
	mongoConfig := GetMongoConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoConfig.Address))

	return client, mongoConfig.DatabaseName, err
}

func MongoConnection() (*mongo.Client, string) {
	once.Do(func() {
		var err error
		newClient, databaseName, err = NewClientDB()
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = newClient.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = newClient.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to database ", err)
		} else {
			log.Println("Connect")
		}
	})

	return newClient, databaseName
}
