package logwrapper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Event struct {
	id      int
	message string
}

var Log *logrus.Entry

func init() {
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

	//Fix for filePath in tests
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Join(filepath.Dir(b), "../../..")
	)
	timeNow := strings.Replace(time.Now().Format(time.UnixDate), ":", "_", -1)
	filePath := fmt.Sprintf("%v/logs/%v.log", basepath, timeNow)
	file, err := os.Create(filePath)
	if err != nil {
		Log.Fatalf("Error creating log file: %v", err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	Log.Logger.SetOutput(writer)
	gin.DefaultWriter = writer
}
