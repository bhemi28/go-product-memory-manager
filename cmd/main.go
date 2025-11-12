package main

import (
	"log"

	"github.com/bhemi28/go-product-memory-manager/cmd/api"
	"github.com/bhemi28/go-product-memory-manager/db"
)

func main() {
	addr := ":8080"
	dbCfg, conn, err := db.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer conn.Close()

	server := api.NewApiServer(addr, dbCfg)

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
