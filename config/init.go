package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init(path string) {
	// Check if loading environment variables from .env file is required
	if os.Getenv("LOAD_CONFIG_FILE") == "" {
		// Load environment variables from .env file
		var err error
		if path != "" {
			err = godotenv.Load(path)
		} else {
			err = godotenv.Load()
		}
		if err != nil {
			log.Fatalf("Error in reading the config file: %v", err)
		}

		//Set gin gonic mode now since gin gonic reads env on init which runs before this function
		mode := os.Getenv(gin.EnvGinMode)
		gin.SetMode(mode)

	}
}
