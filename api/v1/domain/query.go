package domain

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

	"github.com/gin-gonic/gin"
)

func queryDomain(c *gin.Context) {
	db := dbconfig.GetDb()
	var queryRequest GetDomainsQuery
	err := c.BindQuery(&queryRequest)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}
	limit := 10
	offset := (*queryRequest.Page - 1) * limit
	var domains []struct {
		DomainName     string    `json:"domainName"`
		Verified       *bool     `json:"verified"`
		CreatedAt      time.Time `json:"createdAt"`
		Title          string    `json:"title"`
		Headline       string    `json:"headline"`
		Description    string    `json:"description"`
		CoverImageHash string    `json:"coverImageHash"`
		LogoHash       string    `json:"logoHash"`
		Category       string    `json:"category"`
	}

	model := db.Limit(10).Offset(offset).Model(&models.Domain{})
	if queryRequest.Domain != "" {
		model = model.
			Where("domain_name like ?", fmt.Sprintf("%%%s%%", queryRequest.Domain))
	}
	if err := model.
		Where(&models.Domain{Verified: queryRequest.Verified}).
		Select("domain_name, verified, created_at, title, headline, description, cover_image_hash, logo_hash, category").
		Find(&domains).
		Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Error("failed to get domains", err)
		return
	}

	if len(domains) == 0 {
		httpo.NewErrorResponse(200, "No domains found").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Domains fetched successfully", domains).SendD(c)
}
