package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"go-mux-db",
		"5432",
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("database connected")
	}

	return db
}
