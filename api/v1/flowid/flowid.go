package flowid

import (
	"net/http"

	"github.com/TheLazarusNetwork/netsepio-engine/models"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/flowid"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/httphelper"
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
	_, err := hexutil.Decode(walletAddress)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Wallet address (walletAddress) is not valid")
		return
	}
	if walletAddress == "" {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Wallet address (walletAddress) is required")
		return
	}

	flowId, err := flowid.GenerateFlowId(walletAddress, models.AUTH, "")
	if err != nil {
		log.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	userAuthEULA := envutil.MustGetEnv("AUTH_EULA")
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httphelper.SuccessResponse(c, "Flowid successfully generated", payload)
}
