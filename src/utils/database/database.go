package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Close Open Opens the database connection using the enviroment variables
func Open() *sql.DB {
	once.Do(func() {

		dbURI := os.Getenv("DB_URI")
		if dbURI == "" {
			log.Fatal("DB_URI environment variable missing")
		}

		var errOpen error
		db, errOpen = sql.Open("postgres", dbURI)
		if errOpen != nil {
			log.Fatalf("Error opening database: %v", errOpen)
		}

		if errPing := db.Ping(); errPing != nil {
			log.Fatalf("Error connecting to the database: %v", errPing)
		}

		fmt.Println("Database connection open")
	})
	return db
}

// Close Closes the database connection
func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
		fmt.Println("Database connection closed")
	}
}
