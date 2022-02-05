package main

import (
	"github.com/TheLazarusNetwork/netsepio-engine/app"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"
)

func main() {
	app.Init(".env", "logs")
	logwrapper.Log.Info("Starting app")
	port := envutil.MustGetEnv("PORT")
	err := app.GinApp.Run(":" + port)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
