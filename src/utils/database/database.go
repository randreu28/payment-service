package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Open Opens the database connection using the enviroment variables
func Open() *sql.DB {
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Panic("DB_URI environment variable missing")
	}

	db, errOpen := sql.Open("postgres", dbURI)
	if errOpen != nil {
		log.Fatalf("Error opening database: %v", errOpen)
	}

	if errPing := db.Ping(); errPing != nil {
		log.Fatalf("Error connecting to the database: %v", errPing)
	}

	return db
}
