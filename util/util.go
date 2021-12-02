package util

import log "github.com/sirupsen/logrus"

func LogIfError(err error) {
	if err != nil {
		log.Error(err)
	}
}
