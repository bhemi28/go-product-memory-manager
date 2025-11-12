package user

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhemi28/go-product-memory-manager/internal/db"
	internalDB "github.com/bhemi28/go-product-memory-manager/internal/db"
	"github.com/bhemi28/go-product-memory-manager/service/auth"
	"github.com/bhemi28/go-product-memory-manager/types"
	"github.com/bhemi28/go-product-memory-manager/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	db *db.Queries
}

func NewHandler(db *db.Queries) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) chi.Router {
	// Define user-related routes here
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of users"))
	})

	r.Post("/", h.handleUserCreate)
	r.Post("/login", h.handleUserLogin)
	return r
}

func (h *Handler) handleUserCreate(w http.ResponseWriter, r *http.Request) {

	// decoder := json.NewDecoder(r.Body)
	// var req types.CreateUserRequest
	// if err := decoder.Decode(&req); err != nil {
	// 	utils.WriteErrorResponse(w, http.StatusBadRequest, err)
	// }
	var req types.CreateUserRequest
	if err := utils.ParseJson(r, &req); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, err)
	}

	_, err := h.db.GetUserByMail(r.Context(), req.Email)
	if err == nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("the user with mail already exists"))
	}

	hashPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("the password has invalid data"))
	}
	user, err := h.db.CreateUser(r.Context(), internalDB.CreateUserParams{
		ID:       uuid.New(),
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

func (h *Handler) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq types.UserLoginRequest

	if err := utils.ParseJson(r, &loginReq); err != nil {
		log.Print("Invalid req body")
		utils.WriteErrorResponse(w, 400, err)
		return
	}

	_, err := h.db.GetByUserName(r.Context(), loginReq.Username)
	if err != nil {
		utils.WriteErrorResponse(w, 400, err)
	}

	data := map[string]string{
		"username": loginReq.Username,
	}

	accessToken, err := auth.GenerateToken(time.Now().Add(time.Hour*24).Unix(), "halow!", data)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	refreshToken, err := auth.GenerateToken(time.Now().Add(time.Hour*24*30).Unix(), "shalow!", data)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Access token generated:", accessToken)
	log.Println("Refresh token generated:", refreshToken)

	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"success": "true",
		"message": "User logged in successfully",
		"tokens": map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
