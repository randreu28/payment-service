package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Load Loads the enviroment variables
func Load(dir string) {
	err := godotenv.Load(dir)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
