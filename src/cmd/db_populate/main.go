package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	db "payment_service/utils/database"
	"payment_service/utils/env"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	env.Load()
	db := db.Open()
	defer db.Close()
	amount := getAmount()

	// Populate accounts
	for i := 0; i < amount; i++ {
		accountName := fmt.Sprintf("Owner%d", i+1)
		balance := rand.Float64() * 1000
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		insertAccount(db, i+1, accountName, balance, createdAt, createdAt)
		log.Printf("Account with name %s inserted with balance %.2f\n", accountName, balance)
	}

	// Populate transactions
	for i := 0; i < amount; i++ {
		from := rand.Intn(amount) + 1
		to := rand.Intn(amount) + 1
		amount := rand.Float64() * 100
		description := fmt.Sprintf("This is a random generated description with tag %d", i+1)
		createdAt := time.Now().AddDate(0, 0, -rand.Intn(10))
		insertTransaction(db, i+1, from, to, amount, description, createdAt)

		log.Printf("Transaction inserted with amount %.2f from account id %d to account id %d\n", amount, from, to)
	}
}

// insertAccount Inserts a random account to the database
func insertAccount(db *sql.DB, id int, owner string, balance float64, createdAt time.Time, updatedAt time.Time) {
	_, err := db.Exec("INSERT INTO accounts (id, account_owner, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", id, owner, balance, createdAt, updatedAt)
	if err != nil {
		log.Println("Error inserting account:", err)
	}
}

// insertTransaction Inserts a random transaction to the database
func insertTransaction(db *sql.DB, id int, from int, to int, amount float64, description string, createdAt time.Time) {
	_, err := db.Exec("INSERT INTO transactions (id, account_from, account_to, amount, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)", id, from, to, amount, description, createdAt)
	if err != nil {
		log.Println("Error inserting transaction:", err)
	}
}

// getAmount Retrieves the first argument when main is called. If none is provided, it uses a default value of 100
func getAmount() int {
	if len(os.Args) > 1 {
		if num, err := strconv.Atoi(os.Args[1]); err == nil {
			return num
		}
	}
	return 100 // The default value
}
