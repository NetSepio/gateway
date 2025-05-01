package delegatereviewcreation

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/handlers/leaderboard"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
	"netsepio-gateway-v1.1/utils/pkg/aptos"
	"netsepio-gateway-v1.1/utils/pkg/openai"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateReviewCreation")
	{
		g.Use(paseto.PASETO(true))
		g.POST("", deletegateReviewCreation)
	}
}

func deletegateReviewCreation(c *gin.Context) {
	db := database.GetDb()
	var request DelegateReviewCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	walletAddr := c.GetString(paseto.CTX_WALLET_ADDRES)

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

	txResult, err := aptos.DelegateReview(aptos.DelegateReviewParams{Voter: walletAddr, MetaDataUri: request.MetaDataUri, Category: request.Category, DomainAddress: request.DomainAddress, SiteUrl: request.SiteUrl, SiteType: request.SiteType, SiteTag: request.SiteTag, SiteSafety: request.SiteSafety})
	if err != nil {
		if errors.Is(err, aptos.ErrMetadataDuplicated) {
			httpo.NewErrorResponse(http.StatusConflict, "Metadata already exist").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get transaction result", err)
		return
	}
	payload := DelegateReviewCreationPayload{
		TransactionVersion: txResult.Result.Version,
		TransactionHash:    txResult.Result.TransactionHash,
	}

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
		TransactionHash:    txResult.Result.TransactionHash,
		TransactionVersion: txResult.Result.Version,
		SiteRating:         request.SiteRating,
	}
	// go webreview.Publish(request.MetaDataUri, strings.TrimSuffix(request.SiteUrl, "/"))
	if err := db.Create(newReview).Error; err != nil {
		httpo.NewSuccessResponseP(httpo.TXDbFailed, "transaction is successful but failed to store tx in db", payload).Send(c, 200)
		return
	} else {
		userID := c.GetString(paseto.CTX_USER_ID)
		leaderboard.DynamicLeaderBoardUpdate(userID, "reviews")
	}

	httpo.NewSuccessResponseP(200, "request successfully send, review will be delegated soon", payload).SendD(c)
}
