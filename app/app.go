package app

import (
	"netsepio-api/api"
	"netsepio-api/db"
	"os"
	"path/filepath"
	"runtime"

	loggingconfig "netsepio-api/config/logging"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var GinApp *gin.Engine

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
			log.WithFields(loggingconfig.StandardFields).Fatalf("Error in reading the config file: %v", err)
		}
	}
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	db.InitDB()
}
