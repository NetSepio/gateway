package logwrapper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

func Init(basepath string) {
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

	timeNow := strings.Replace(time.Now().Format(time.UnixDate), ":", "_", -1)
	fileName := fmt.Sprintf("%v.log", timeNow)
	filePath := filepath.Join(basepath, fileName)
	logToFile, ok := os.LookupEnv("LOG_TO_FILE")
	if !ok {
		Log.Fatal("env var LOG_TO_FILE is undefined")
	}
	if logToFile == "true" {
		file, err := os.Create(filePath)
		if err != nil {
			Log.Fatalf("Error creating log file: %v", err)
		}
		writer := io.MultiWriter(file, os.Stdout)
		Log.Logger.SetOutput(writer)
		gin.DefaultWriter = writer
	}
}
