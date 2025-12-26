package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	files := []string{
		".env",        // base
		".env." + env, // env-specific
	}

	for _, f := range files {
		if err := godotenv.Overload(f); err == nil {
			log.Printf("Loaded env file: %s", f)
		}
	}
}
