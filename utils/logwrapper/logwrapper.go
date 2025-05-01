package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
	"netsepio-gateway-v1.1/utils/load"
)

var Log *logrus.Entry

func Init() {
	appName := load.Cfg.APP_NAME
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
