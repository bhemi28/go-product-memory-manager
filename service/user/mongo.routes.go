package user

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoHandler struct {
	db *mongo.Client
}

func NewMongoUserHandler(db *mongo.Client) *mongoHandler {
	return &mongoHandler{
		db: db,
	}
}

func (m *mongoHandler) RegisterRoutes(r chi.Router) {
	r.Get('/users', m.)
}

func 
