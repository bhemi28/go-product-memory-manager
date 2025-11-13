package user

import (
	"fmt"
	"net/http"

	mongodb "github.com/bhemi28/go-product-memory-manager/mongo/db"
	"github.com/bhemi28/go-product-memory-manager/service/auth"
	"github.com/bhemi28/go-product-memory-manager/types"
	"github.com/bhemi28/go-product-memory-manager/utils"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoHandler struct {
	userClient *mongodb.UserClient
}

func NewMongoUserHandler(db *mongo.Client) *mongoHandler {
	return &mongoHandler{
		userClient: mongodb.NewUserClient(db),
	}
}

func (m *mongoHandler) RegisterRoutes(r chi.Router) chi.Router {
	r.Post("/", m.createUser)

	return r
}

func (m *mongoHandler) createUser(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var req types.CreateUserRequest
	// if err := decoder.Decode(&req); err != nil {
	// 	utils.WriteErrorResponse(w, http.StatusBadRequest, err)
	// }
	var req types.CreateUserRequest
	if err := utils.ParseJson(r, &req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, err)
	}

	// _, err := h.db.GetUserByMail(r.Context(), req.Email)
	// if err == nil {
	// 	utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("the user with mail already exists"))
	// }

	hashPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("the password has invalid data"))
	}
	user, err := m.userClient.CreateUser(r.Context(), mongodb.UserSchema{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	})

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]any{
		"success": "true",
		"user":    user,
		"message": "User created successfully",
	})
}
