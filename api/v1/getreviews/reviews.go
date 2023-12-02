package getreviews

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"gorm.io/gorm"

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
	limit := 12
	offset := (*queryRequest.Page - 1) * limit
	var reviews []struct {
		Voter              string `json:"voter"`
		Name               string `json:"name"`
		MetaDataUri        string `json:"metaDataUri"`
		Category           string `json:"category"`
		DomainAddress      string `json:"domainAddress"`
		SiteUrl            string `json:"siteUrl"`
		SiteType           string `json:"siteType"`
		SiteTag            string `json:"siteTag"`
		SiteSafety         string `json:"siteSafety"`
		SiteIpfsHash       string `json:"siteIpfsHash"`
		TransactionHash    string `json:"transactionHash"`
		TransactionVersion int64  `json:"transactionVersion"`
		DeletedAt          gorm.DeletedAt
		CreatedAt          time.Time `json:"createdAt"`
	}

	if err := db.Limit(10).Offset(offset).Joins("left join users ON reviews.voter = users.wallet_address").Model(&models.Review{}).Order("reviews.created_at desc").
		Where(&models.Review{Voter: strings.ToLower(queryRequest.Voter), DomainAddress: queryRequest.Domain}).
		Select("reviews.meta_data_uri, reviews.category, reviews.domain_address, reviews.site_url, reviews.site_type, reviews.site_tag, reviews.site_safety, reviews.site_ipfs_hash, reviews.transaction_hash, reviews.transaction_version, reviews.created_at, reviews.voter, users.name").
		Find(&reviews).
		Error; err != nil {
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
			Voter:              reviews[i].Voter,
			Name:               reviews[i].Name,
		}
	}

	if len(payload) == 0 {
		httpo.NewErrorResponse(404, "No reviews found").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Reviews fetched successfully", payload).SendD(c)
}
