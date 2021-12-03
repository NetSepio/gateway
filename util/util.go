package util

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func LogIfError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func TFatalIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
