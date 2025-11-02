package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables instead")
	}

	requiredVars := []string{"DB_URL", "JWT_KEY"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Printf("Warning: Environment variable %s is not set\n", v)
		}
	}

	fmt.Println("Environment variables loaded")
}
