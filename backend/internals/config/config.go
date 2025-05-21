package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"errors" // Added missing import for errors
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return errors.New("failed to load environment variables") // Corrected error creation
	}
	return nil
}