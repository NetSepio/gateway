package delegatereviewcreation

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/smartcontract/rawtrasaction"
	"github.com/NetSepio/gateway/generated/smartcontract/gennetsepio"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateReviewCreation")
	{
		g.Use(paseto.PASETO)
		g.POST("", deletegateReviewCreation)
	}
}

func deletegateReviewCreation(c *gin.Context) {
	var request DelegateReviewCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	voterAddr := common.HexToAddress(request.Voter)
	abiS := gennetsepio.GennetsepioABI

	tx, err := rawtrasaction.SendRawTrasac(abiS, "delegateReviewCreation", request.Category, request.DomainAddress, request.SiteUrl, request.SiteType, request.SiteTag, request.SiteSafety, request.MetaDataUri, voterAddr)

	if err != nil {
		httphelper.NewInternalServerError(c, "failed to call %v of %v, error: %v", "delegateReviewCreation", "NETSEPIO", err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	payload := DelegateReviewCreationPayload{
		TransactionHash: transactionHash,
	}
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	httphelper.SuccessResponse(c, "request successfully send, review will be delegated soon", payload)
}
