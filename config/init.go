package config

import (
	"os"
	"path/filepath"
	"runtime"

	. "netsepio-api/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	// Check if loading environment variables from .env file is required
	if os.Getenv("LOAD_CONFIG_FILE") == "" {
		// Load environment variables from .env file
		var (
			_, b, _, _ = runtime.Caller(0)
			basepath   = filepath.Dir(filepath.Dir(b))
		)

		err := godotenv.Load(filepath.Join(basepath, ".env"))
		if err != nil {
			Log.Fatalf("Error in reading the config file: %v", err)
		}

		//Set gin gonic mode now since gin gonic reads env on init which runs before this function
		mode := os.Getenv(gin.EnvGinMode)
		gin.SetMode(mode)

	}
}
