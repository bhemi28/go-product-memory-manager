package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	Username  string              `bson:"username"`
	Email     string              `bson:"email"`
	Password  string              `bson:"password"`
	IsDeleted *bool               `bson:"isDeleted"`
	CreatedAt primitive.DateTime  `bson:"createdAt"`
	UpdatedAt primitive.DateTime  `bson:"updatedAt"`
}
