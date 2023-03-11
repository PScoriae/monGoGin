package db

import (
	"context"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:3500/mongogin-db"

var MongoClient *mongo.Client
var MongoClientError error
var mongoOnce sync.Once

func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
		var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.TODO(), opts)

		MongoClient = client
		MongoClientError = err
	})
	return MongoClient, MongoClientError
}
