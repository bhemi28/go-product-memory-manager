package main

import (
	"context"
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
	mfg := db.ConnectToMongo()

	defer conn.Close()
	defer mfg.Disconnect(context.Background())

	server := api.NewApiServer(addr, dbCfg, mfg)

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
