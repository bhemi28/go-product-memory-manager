package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	uri := "mongodb://172.24.80.1:27017/db"

	connection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := connection.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return connection
}
