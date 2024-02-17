package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load loads the environment variables from the .env.local file located relative to the path of the Go project
func Load() {
	executablePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting executable path: %v", err)
	}

	basePath := filepath.Dir(executablePath)
	filePath := filepath.Join(basePath, ".env.local")

	err = godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
	}

	log.Printf("Enviroment variables loaded")
}
