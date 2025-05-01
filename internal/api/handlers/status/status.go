package status

import (
	"crypto/ed25519"
	"encoding/hex"
	"runtime"

	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/utils/load"
)

var pubKey = ""

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
}

func status(c *gin.Context) {

	k := load.Cfg.PASETO_PRIVATE_KEY
	pvKey := ed25519.PrivateKey(k)
	pubKey = "0x" + hex.EncodeToString(pvKey.Public().(ed25519.PublicKey))

	c.JSON(200, gin.H{
		"status":    "alive",
		"publicKey": pubKey,
		"goVersion": runtime.Version(),
		"version":   load.Cfg.VERSION,
		"network":   load.Cfg.NETWORK,
	})
}
