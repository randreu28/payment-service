package main

import (
	"log"
	db "payment_service/utils/database"
	"payment_service/utils/env"
)

func main() {
	env.Load()
	db := db.Open()
	defer db.Close()
	_, err := db.Exec("DELETE FROM accounts")
	if err != nil {
		log.Println("Error deleting accounts:", err)
	}

	_, err = db.Exec("DELETE FROM transactions")
	if err != nil {
		log.Println("Error deleting transactions:", err)
	}
}
