package authenticate

import (
	"net/http"
	"netsepio-api/models/claims"
	"netsepio-api/util/pkg/auth"
	"netsepio-api/util/pkg/cryptosign"
	"os"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {

	//TODO remove flow id if 200
	var req AuthenticateRequest
	c.BindJSON(&req)

	// Append userId to the message
	message := req.FlowId + "m"
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
		jwtToken, err := auth.GenerateToken(customClaims, jwtPrivateKey)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error Occured")
		}
		c.JSON(http.StatusOK, map[string]string{
			"token": jwtToken,
		})
	} else {
		c.String(http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}
