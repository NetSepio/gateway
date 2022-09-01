package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

func Init() {
	appName, ok := os.LookupEnv("APP_NAME")

	if !ok {
		appName = "web3-auth"
	}
	hostname, err := os.Hostname()

	Log = logrus.New().WithFields(logrus.Fields{
		"hostname": hostname,
		"appname":  appName,
	})
	if err != nil {
		Log.Warnf("Error in getting hostname: %v", err)
	}
	Log.Logger.SetFormatter(&logrus.JSONFormatter{})
}
