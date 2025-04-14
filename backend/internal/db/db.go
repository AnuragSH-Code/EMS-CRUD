package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetMongoClient() *mongo.Client {
	return client
}

func InitDB(uri string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)

	var err error

	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Database Connect")

	return nil

}

func GetConnection(database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func CloseDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client.Disconnect(ctx)
		fmt.Println("Database connection closed")
	}
}
