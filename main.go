package main

import (
	"github.com/NetSepio/gateway/app"
	"github.com/NetSepio/gateway/util/pkg/envutil"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
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
