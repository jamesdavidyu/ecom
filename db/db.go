package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error) {
	godotenv.Load()

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
