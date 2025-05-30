package flowid

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/pkg/flowid"

	"github.com/gin-gonic/gin"
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
			load.Logger.Error(err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	} else {
		var err error
		flowId, err = flowid.GenerateFlowId(walletAddress, models.AUTH, "", userId)
		if err != nil {
			load.Logger.Error(err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	}
	userAuthEULA := load.Cfg.AUTH_EULA
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.NewSuccessResponseP(200, "Flowid successfully generated", payload).SendD(c)
}
