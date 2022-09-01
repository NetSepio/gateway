package main

import (
	"fmt"

	"github.com/NetSepio/gateway/app"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting app")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
