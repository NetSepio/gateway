package main

import (
	"netsepio-api/app"
	"netsepio-api/util/pkg/logwrapper"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting app")
	err := app.GinApp.Run(":8000")
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
