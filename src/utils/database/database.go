package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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

// PopulateDB populates the database with random accounts and transactions
func PopulateDB(db *sql.DB, amount int) {

	// Populate accounts
	for i := 0; i < amount; i++ {
		accountName := fmt.Sprintf("Owner%d", i+1)
		balance := rand.Float64() * 1000
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		_, err := db.Exec("INSERT INTO accounts (id, account_owner, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", i+1, accountName, balance, createdAt, createdAt)
		if err != nil {
			log.Println("Error inserting account:", err)
			return
		}

		log.Printf("Account with name %s inserted with balance %.2f\n", accountName, balance)
	}

	// Populate transactions
	for i := 0; i < amount; i++ {
		from := rand.Intn(amount) + 1
		to := rand.Intn(amount) + 1
		amount := rand.Float64() * 100
		description := fmt.Sprintf("This is a random generated description with tag %d", i+1)
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		_, err := db.Exec("INSERT INTO transactions (id, account_from, account_to, amount, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)", i+1, from, to, amount, description, createdAt)
		if err != nil {
			log.Println("Error inserting transaction:", err)
			return
		}

		log.Printf("Transaction inserted with amount %.2f from account id %d to account id %d\n", amount, from, to)
	}
}

func CleanupDB(db *sql.DB) {
	_, err := db.Exec("DELETE FROM accounts")
	if err != nil {
		log.Println("Error deleting accounts:", err)
		return
	}

	_, err = db.Exec("DELETE FROM transactions")
	if err != nil {
		log.Println("Error deleting transactions:", err)
		return
	}

	log.Println("Database cleanup completed successfully.")
}
