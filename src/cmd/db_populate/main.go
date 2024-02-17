package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Loads enviroment variables
	err := godotenv.Load(".env.local")
	dbURI := os.Getenv("DB_URI")

	if err != nil || dbURI == "" {
		panic("Enviroment variables missing")
	}

	// Opens a TCP/IP connection to the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database")
	amount := getAmount()

	// Populate accounts
	for i := 0; i < amount; i++ {
		accountName := fmt.Sprintf("Owner%d", i+1)
		balance := rand.Float64() * 1000
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		insertAccount(db, accountName, balance, createdAt, createdAt)
		log.Printf("Account with name %s inserted with balance %.2f\n", accountName, balance)
	}

	// Populate transactions
	for i := 0; i < amount; i++ {
		from := rand.Intn(10) + 1
		to := rand.Intn(10) + 1
		amount := rand.Float64() * 100
		description := fmt.Sprintf("This is a random generated description with tag %d", i+1)
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		insertTransaction(db, from, to, amount, description, createdAt)

		log.Printf("Transaction inserted with amount %.2f from account id %d to account id %d\n", amount, from, to)
	}
}

// insertAccount Inserts a random account to the database
func insertAccount(db *sql.DB, owner string, balance float64, createdAt time.Time, updatedAt time.Time) {
	_, err := db.Exec("INSERT INTO accounts (account_owner, balance, created_at, updated_at) VALUES ($1, $2, $3, $4)", owner, balance, createdAt, updatedAt)
	if err != nil {
		log.Println("Error inserting account:", err)
	}
}

// insertTransaction Inserts a random transaction to the database
func insertTransaction(db *sql.DB, from, to int, amount float64, description string, createdAt time.Time) {
	_, err := db.Exec("INSERT INTO transactions (account_from, account_to, amount, description, created_at) VALUES ($1, $2, $3, $4, $5)", from, to, amount, description, createdAt)
	if err != nil {
		log.Println("Error inserting transaction:", err)
	}
}

// getAmount Retrieves the first argument when main is called. If none is provided, it uses a default value
func getAmount() int {
	if len(os.Args) > 1 {
		if num, err := strconv.Atoi(os.Args[1]); err == nil {
			return num
		}
	}
	return 100 // The default value
}
