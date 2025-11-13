package api

import (
	"log"
	"net/http"

	"github.com/bhemi28/go-product-memory-manager/internal/db"
	"github.com/bhemi28/go-product-memory-manager/service/user"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIserver struct {
	addr string
	cfg  *db.Queries
	mfg *mongo.Client
}

func NewApiServer(addr string, cfg *db.Queries, mfg *mongo.Client) *APIserver {
	return &APIserver{
		addr: addr,
		cfg:  cfg,
		mfg:  mfg,
	}
}

func (s *APIserver) Start() error {
	router := chi.NewRouter()
	// Define your routes here
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		// r.Mount("/users", user.NewHandler(s.cfg).RegisterRoutes(chi.NewRouter()))
		r.Mount("/users", user.NewMongoUserHandler(s.mfg).RegisterRoutes(chi.NewRouter()))

		log.Println("Starting server on", s.addr)
	})
	return http.ListenAndServe(s.addr, router)
}
