package logginconfig

import log "github.com/sirupsen/logrus"

var StandardFields = log.Fields{
	"hostname": "HostServer",
	"appname":  "BCMWallet",
}
