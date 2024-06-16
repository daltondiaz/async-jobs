package conf

import (
	"log"

	"github.com/joho/godotenv"
)

// Load the Environment variables
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		log.Fatal("Error to load .env file on root of project")
	}
}
