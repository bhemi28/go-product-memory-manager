package api

import (
	"log"
	"net/http"

	"github.com/bhemi28/go-product-memory-manager/internal/db"
	"github.com/bhemi28/go-product-memory-manager/service/user"
	"github.com/go-chi/chi"
)

type APIserver struct {
	addr string
	cfg  *db.Queries
}

func NewApiServer(addr string, cfg *db.Queries) *APIserver {
	return &APIserver{
		addr: addr,
		cfg:  cfg,
	}
}

func (s *APIserver) Start() error {
	router := chi.NewRouter()
	// Define your routes here
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		r.Mount("/users", user.NewHandler(s.cfg).RegisterRoutes(chi.NewRouter()))

		log.Println("Starting server on", s.addr)
	})
	return http.ListenAndServe(s.addr, router)
}
