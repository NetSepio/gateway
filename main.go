package main

import (
	"github.com/TheLazarusNetwork/netsepio-engine/app"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"
)

func main() {
	app.Init(".env", "logs")
	logwrapper.Log.Info("Starting app")
	err := app.GinApp.Run(":8000")
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
