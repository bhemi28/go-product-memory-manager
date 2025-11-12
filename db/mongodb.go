package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")

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
