package main

import (
	"database/sql"
	"log"
	"net/http"

	handler "github.com/jamesdavidyu/ecom/api"
	"github.com/jamesdavidyu/ecom/cmd/api"
	"github.com/jamesdavidyu/ecom/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api", handler.Handler)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
