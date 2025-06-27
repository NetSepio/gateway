package deletereview

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/NetSepio/gateway/utils/pkg/aptos"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/deleteReview")
	{
		g.Use(paseto.PASETO(false))
		g.DELETE("", deleteReview)
	}
}

func deleteReview(c *gin.Context) {
	db := database.GetDb()
	var request DeleteReviewRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, "payload is invalid").SendD(c)
		return
	}
	walletAddr := c.GetString(paseto.CTX_WALLET_ADDRES)
	var review models.Review
	if err = db.First(&review, models.Review{Voter: walletAddr, MetaDataUri: request.MetaDataUri}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "review not found").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get review details", err)
		return
	}

	txResult, err := aptos.DeleteReview(request.MetaDataUri)
	if err != nil {
		if errors.Is(err, aptos.ErrMetadataNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "Metadata not found").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get transaction result", err)
		return
	}
	payload := DeleteReviewPayload{
		TransactionVersion: txResult.Result.Version,
		TransactionHash:    txResult.Result.TransactionHash,
	}

	if err := db.Delete(&models.Review{}, "meta_data_uri = ?", request.MetaDataUri).Error; err != nil {
		httpo.NewSuccessResponseP(httpo.TXDbFailed, "transaction is successful but failed to delete review from db", payload).Send(c, 200)
		return
	}

	httpo.NewSuccessResponseP(200, "request successfully send, review will be deleted soon", payload).SendD(c)
}
