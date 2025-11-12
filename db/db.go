package db

import (
	"database/sql"
	"log"

	"github.com/bhemi28/go-product-memory-manager/internal/db" // Import your db package

	_ "github.com/lib/pq" // PostgreSQL driver
)

type ApiConfig struct {
	DB *db.Queries
}

func ConnectToDB() (*db.Queries, *sql.DB, error) {
	connectionString := "postgres://bhemi:PSQL@bhemi@localhost:5432/product_manager?sslmode=disable"
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, conn, err
	}

	cfg := ApiConfig{
		DB: db.New(conn),
	}

	return cfg.DB, conn, nil

}
