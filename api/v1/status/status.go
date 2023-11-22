package status

import (
	"crypto/ed25519"
	"encoding/hex"
	"runtime"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
)

var pubKey = ""

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
	k := envconfig.EnvVars.PASETO_PRIVATE_KEY
	pvKey := ed25519.PrivateKey(k)
	pubKey = "0x" + hex.EncodeToString(pvKey.Public().(ed25519.PublicKey))
}

func status(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":    "alive",
		"publicKey": pubKey,
		"goVersion": runtime.Version(),
		"version":   envconfig.EnvVars.VERSION,
		"network":   envconfig.EnvVars.NETWORK,
	})
}
