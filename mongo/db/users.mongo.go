package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserClient struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserClient(conn *mongo.Client) *UserClient {
	return &UserClient{
		client: conn,
		coll:   conn.Database("db").Collection("users"),
	}
}

func (u *UserClient) CreateUser(ctx context.Context, body UserSchema) (*mongo.InsertOneResult, error) {
	res, err := u.coll.InsertOne(ctx, body)

	if err != nil {
		// Don’t crash the server; propagate the real error up
		log.Printf("mongodb insert user failed: %v", err)
		return nil, fmt.Errorf("mongodb insert user failed: %w", err)
	}

	return res, nil
}
