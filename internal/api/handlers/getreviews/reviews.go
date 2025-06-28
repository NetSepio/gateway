package getreviews

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"

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
	db := database.GetDb()
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
		SiteRating         int    `json:"siteRating"`
		TransactionHash    string `json:"transactionHash"`
		TransactionVersion int64  `json:"transactionVersion"`
		DeletedAt          gorm.DeletedAt
		CreatedAt          time.Time `json:"createdAt"`
	}

	if err := db.Limit(limit).Offset(offset).Joins("left join users ON reviews.voter = users.wallet_address").Model(&models.Review{}).Order("reviews.created_at desc").
		Where(&models.Review{Voter: strings.ToLower(queryRequest.Voter), DomainAddress: queryRequest.Domain}).
		Select("reviews.meta_data_uri, reviews.category, reviews.domain_address, reviews.site_url, reviews.site_type, reviews.site_tag, reviews.site_safety, reviews.site_ipfs_hash, reviews.transaction_hash, reviews.transaction_version, reviews.created_at, reviews.voter, reviews.site_rating, users.name").
		Find(&reviews).
		Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get reviews", err)
		return
	}
	var averageRating *float64 = nil
	var totalReviews int64
	if queryRequest.Domain == "" {
		// get total reviews
		err = db.Model(&models.Review{}).Count(&totalReviews).Error
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Error("failed to get reviews", err)
			return
		}
	} else {
		// Query to calculate average rating for the site URL
		err = db.Model(&models.Review{}).Where("domain_address = ?", queryRequest.Domain).Select("AVG(site_rating)").Row().Scan(&averageRating)
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Error("failed to get reviews", err)
			return
		}

		// Get total reviews for the site URL
		err = db.Model(&models.Review{}).Where("domain_address = ?", queryRequest.Domain).Count(&totalReviews).Error
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Error("failed to get reviews", err)
			return
		}

	}
	var reviewsPayload []GetReviewsItem = make([]GetReviewsItem, len(reviews))
	for i := 0; i < len(reviews); i++ {
		reviewsPayload[i] = GetReviewsItem{
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
			SiteRating:         reviews[i].SiteRating,
		}
	}

	payload := GetReviewsPayload{
		Reviews:       reviewsPayload,
		TotalReviews:  totalReviews,
		AverageRating: averageRating,
	}

	if len(reviewsPayload) == 0 {
		httpo.NewErrorResponse(404, "No reviews found").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Reviews fetched successfully", payload).SendD(c)
}
