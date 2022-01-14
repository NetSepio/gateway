package main

import (
	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting app")
	err := app.GinApp.Run(":8000")
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
