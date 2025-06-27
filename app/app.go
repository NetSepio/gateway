package app

import (
	"netsepio-gateway-v1.1/internal/caching"
	"netsepio-gateway-v1.1/internal/database"
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
	database.Migrate()

	// Initialize Redis
	caching.InitRedis()

	// Initialize the of Goroutines
	// routines.Init()

	// Initialize the P2P node
	// p2pnode.Init()

	// Initialize the server
	server.Init()

	load.Logger.Sugar().Infoln("App initialized successfully.")

}
