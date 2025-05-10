package flowid

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/flowid"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.Use(paseto.PASETO(true))
		g.GET("", GetFlowId)
	}
}

func GetFlowId(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	walletAddress := c.Query("walletAddress")
	chain_symbol := c.Query("chain")

	if walletAddress == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "Wallet address (walletAddress) is required").SendD(c)
		return
	}
	if chain_symbol != "sol" {
		_, err := hexutil.Decode(walletAddress)
		if err != nil {
			httpo.NewErrorResponse(http.StatusBadRequest, "Please pass the valid chain name").SendD(c)
			return
		}
	}
	var flowId string
	if chain_symbol == "sol" {
		var err error
		flowId, err = flowid.GenerateFlowIdSol(walletAddress, models.AUTH, "", userId)
		if err != nil {
			log.Error(err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	} else {
		var err error
		flowId, err = flowid.GenerateFlowId(walletAddress, models.AUTH, "", userId)
		if err != nil {
			log.Error(err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)

			return
		}
	}
	userAuthEULA := envconfig.EnvVars.AUTH_EULA
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.NewSuccessResponseP(200, "Flowid successfully generated", payload).SendD(c)
}
