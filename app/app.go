package app

import (
	"github.com/NetSepio/gateway/internal/caching"
	"github.com/NetSepio/gateway/internal/database"
	p2pnode "github.com/NetSepio/gateway/internal/p2p-Node"
	"github.com/NetSepio/gateway/internal/routines"
	"github.com/NetSepio/gateway/internal/server"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
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
	routines.Init()

	// // Initialize the P2P node
	p2pnode.Init()

	// Initialize the server
	server.Init()

	load.Logger.Sugar().Infoln("App initialized successfully.")

}
