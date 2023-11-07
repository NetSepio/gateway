package flowid

import (
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/flowid"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.GET("", GetFlowId)
	}
}

func GetFlowId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "Wallet address (walletAddress) is required").SendD(c)
		return
	}
	_, err := hexutil.Decode(walletAddress)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "Wallet address (walletAddress) is not valid").SendD(c)
		return
	}
	flowId, err := flowid.GenerateFlowId(walletAddress, models.AUTH, "")
	if err != nil {
		log.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)

		return
	}
	userAuthEULA := envconfig.EnvVars.AUTH_EULA
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.NewSuccessResponseP(200, "Flowid successfully generated", payload).SendD(c)
}
