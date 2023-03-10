package db

import "go.mongodb.org/mongo-driver/mongo"

const database string = "mongogin-prod"

func GetUserColl() (*mongo.Collection, error) {
	client, err := GetMongoClient()
	return client.Database(database).Collection("users"), err
}
