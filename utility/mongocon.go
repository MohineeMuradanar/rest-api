package utility

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client ...
var Client *mongo.Client

func MongoConnection() {

	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://root:root@mongo-service:27017")

	Client, _ = mongo.Connect(ctx, clientOptions)

}

func DB() *mongo.Collection {
	MongoConnection()
	collection := Client.Database("University").Collection("studentdata")
	return collection
}
