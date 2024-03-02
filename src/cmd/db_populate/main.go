package main

import (
	"os"
	db "payment_service/utils/database"
	"payment_service/utils/env"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	env.Load()
	database := db.Open()
	defer database.Close()
	amount := getAmount()

	db.PopulateDB(database, amount)
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
