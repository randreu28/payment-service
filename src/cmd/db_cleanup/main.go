package main

import (
	db "payment_service/utils/database"
	"payment_service/utils/env"
)

func main() {
	env.Load()
	database := db.Open()
	defer database.Close()
	db.CleanupDB(database)
}
