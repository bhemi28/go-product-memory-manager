package product

import (
	"net/http"

	"github.com/bhemi28/go-product-memory-manager/internal/db"
	"github.com/bhemi28/go-product-memory-manager/types"
	"github.com/bhemi28/go-product-memory-manager/utils"
	"github.com/go-chi/chi"
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
	// Define product-related routes here
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of products"))
	})

	r.Post("/", h.handleProductCreate)
	// r.Get("/{id}", h.handleProductGet)
	return r
}

func (h *Handler) handleProductCreate(w http.ResponseWriter, r *http.Request) {
	var req types.ProductCreate
	if err := utils.ParseJson(r, &req); err != nil {
		utils.WriteErrorResponse(w, 400, err)
		return
	}

	utils.WriteJSONResponse(w, 201, nil)
	return
}
