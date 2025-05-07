package app

import (
	"netsepio-gateway-v1.1/internal/caching"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/internal/routines"
	"netsepio-gateway-v1.1/internal/server"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

// Initialize the app
func Init() {

	load.Logger.Sugar().Infoln("Initializing the app...")
	// Initialize the logger
	logwrapper.Init()
	// test db connection
	database.GetDb()

	// Migrate the database
	// database.Migrate()
	
	// Initialize Redis
	caching.InitRedis()

	// Initialize the server
	server.Init()

	// Initialize the of Goroutines
	routines.Init()

	load.Logger.Sugar().Infoln("App initialized successfully.")

}
