package main

import (
	"database/sql"
	"log"

	handler "github.com/jamesdavidyu/ecom/api"
	"github.com/jamesdavidyu/ecom/db"
)

func main() {
	handler.ExportedFunction()

	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := handler.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
