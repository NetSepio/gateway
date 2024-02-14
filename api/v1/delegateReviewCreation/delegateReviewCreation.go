package delegatereviewcreation

import (
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/smartcontract/rawtrasaction"
	"github.com/NetSepio/gateway/generated/smartcontract/gennetsepio"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/pkg/openai"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateReviewCreation")
	{
		g.Use(paseto.PASETO(false))
		g.POST("", deletegateReviewCreation)
	}
}

func deletegateReviewCreation(c *gin.Context) {
	db := dbconfig.GetDb()

	var request DelegateReviewCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	walletAddr := c.GetString(paseto.CTX_WALLET_ADDRES)
	voterAddr := common.HexToAddress(walletAddr)
	abiS := gennetsepio.GennetsepioABI

	isSpam, err := openai.IsReviewSpam(request.Description)
	if err != nil {
		logwrapper.Error("failed to check spam", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	if isSpam {
		httpo.NewErrorResponse(http.StatusForbidden, "Review is spam").SendD(c)
		return
	}
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

	newReview := &models.Review{
		Voter:              walletAddr,
		MetaDataUri:        request.MetaDataUri,
		Category:           request.Category,
		DomainAddress:      request.DomainAddress,
		SiteUrl:            strings.TrimSuffix(request.SiteUrl, "/"),
		SiteType:           request.SiteType,
		SiteTag:            request.SiteTag,
		SiteSafety:         request.SiteSafety,
		SiteIpfsHash:       "",
		TransactionHash:    transactionHash,
		TransactionVersion: 0,
		SiteRating:         request.SiteRating,
	}
	// go webreview.Publish(request.MetaDataUri, strings.TrimSuffix(request.SiteUrl, "/"))
	if err := db.Create(newReview).Error; err != nil {
		httpo.NewSuccessResponseP(httpo.TXDbFailed, "transaction is successful but failed to store tx in db", payload).Send(c, 200)
		return
	}
	httphelper.SuccessResponse(c, "request successfully send, review will be delegated soon", payload)
}
