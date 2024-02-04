package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Load Loads the enviroment variables
func Load() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
