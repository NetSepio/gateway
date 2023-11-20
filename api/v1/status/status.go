package status

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
}

func status(c *gin.Context) {
	k := envconfig.EnvVars.PASETO_PRIVATE_KEY
	ed25519_pv_key := ed25519.PrivateKey(k)
	c.JSON(200, gin.H{
		"status":    "alive",
		"publicKey": "0x" + hex.EncodeToString(ed25519_pv_key.Public().(ed25519.PublicKey)),
	})
}
