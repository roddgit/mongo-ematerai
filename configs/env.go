package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	return os.Getenv("MONGOURI")
}

func EnvMaterai() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	// result := make(map[string][]string)
	result := make(map[string]string)

	// result["BASE_API"] = append(result["BASE_API"], os.Getenv("BASE_API"))
	result["BASE_API"] = os.Getenv("BASE_API")
	result["API_LOGIN"] = os.Getenv("API_LOGIN")

	return result
}
