package delegateartifactcreation

import (
	"net/http"

	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract/rawtrasaction"
	gcreatify "github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateArtifactCreation")
	{
		g.POST("", deletegateArtifactCreation)
	}
}

func deletegateArtifactCreation(c *gin.Context) {
	var request DelegateArtifactCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	creatorAddr := common.HexToAddress(request.CreatorAddress)
	abiS := gcreatify.CreatifyABI

	tx, err := rawtrasaction.SendRawTrasac(abiS, "delegateArtifactCreation", creatorAddr, request.MetaDataHash)

	if err != nil {
		httphelper.NewInternalServerError(c, "failed to call %v of %v, error: %v", "delegateArtifactCreation", "Creatify", err.Error())
		return
	}
	logwrapper.Infof("trasaction hash is %v", tx.Hash().String())
}
