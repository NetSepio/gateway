package getreviews

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/getreviews")
	{
		g.GET("", getReviews)
	}
}

func getReviews(c *gin.Context) {
	db := dbconfig.GetDb()
	var queryRequest GetReviewsQuery
	err := c.BindQuery(&queryRequest)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}
	walletAddr := c.Query("walletAddress")
	limit := 10
	offset := (*queryRequest.Page - 1) * limit
	var reviews []models.Review
	if err := db.Limit(10).Offset(offset).Find(&reviews, models.Review{Voter: walletAddr}).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get reviews", err)
		return
	}

	var payload GetReviewsPayload = make(GetReviewsPayload, len(reviews))
	for i := 0; i < len(reviews); i++ {
		payload[i] = GetReviewsItem{
			MetaDataUri:        reviews[i].MetaDataUri,
			Category:           reviews[i].Category,
			DomainAddress:      reviews[i].DomainAddress,
			SiteUrl:            reviews[i].SiteUrl,
			SiteType:           reviews[i].SiteType,
			SiteTag:            reviews[i].SiteTag,
			SiteSafety:         reviews[i].SiteSafety,
			SiteIpfsHash:       reviews[i].SiteIpfsHash,
			TransactionHash:    reviews[i].TransactionHash,
			TransactionVersion: reviews[i].TransactionVersion,
			CreatedAt:          reviews[i].CreatedAt,
		}
	}

	if len(payload) == 0 {
		httpo.NewSuccessResponseP(200, "No reviews found", payload).SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Reviews fetched successfully", payload).SendD(c)
}
