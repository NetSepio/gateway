package domain

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
)

func queryDomain(c *gin.Context) {
	db := database.GetDb()
	var queryRequest GetDomainsQuery
	err := c.BindQuery(&queryRequest)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}
	limit := 12
	offset := (*queryRequest.Page - 1) * limit
	var domains []struct {
		Id             string    `json:"id"`
		DomainName     string    `json:"domainName"`
		Verified       *bool     `json:"verified"`
		CreatedAt      time.Time `json:"createdAt"`
		Title          string    `json:"title"`
		Headline       string    `json:"headline"`
		Description    string    `json:"description"`
		CoverImageHash string    `json:"coverImageHash"`
		LogoHash       string    `json:"logoHash"`
		Category       string    `json:"category"`
		Blockchain     string    `json:"blockchain"`
		CreatedBy      string    `json:"createdBy"`
		CreatorName    string    `json:"creatorName"`
		Claimable      string    `json:"claimable"`
		TxtValue       string    `json:"txtValue,omitempty"`
	}

	model := db.Limit(limit).Offset(offset).Model(&models.Domain{}).Order("title ASC")
	if queryRequest.DomainName != "" {
		model = model.
			Where("domain_name like ?", fmt.Sprintf("%%%s%%", queryRequest.DomainName)).Where(&models.Domain{Id: queryRequest.DomainId})
	}

	if queryRequest.VerifiedWithClaimable {
		model = model.
			Where("verified = true or claimable = true")
	} else {
		model = model.
			Where(&models.Domain{Verified: queryRequest.Verified})
	}

	if queryRequest.OnlyAdmin {
		userId := c.GetString(paseto.CTX_USER_ID)
		if userId == "" {
			httpo.NewErrorResponse(http.StatusBadRequest, "auth token required if onlyAdmin is true").SendD(c)
			return
		}
		if err := model.Where("da.admin_id = ?", userId).
			Select("id, domain_name, verified, created_at, title, headline, description, cover_image_hash, logo_hash, category, blockchain, created_by_id created_by, u.name creator_name, txt_value, claimable").
			Joins("INNER JOIN users u ON u.user_id = created_by_id").
			Joins("INNER JOIN domain_admins da ON da.domain_id = domains.id").
			Find(&domains).
			Error; err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Error("failed to get domains", err)
			return
		}
	} else {
		if err := model.
			Select("id, domain_name, verified, created_at, title, headline, description, cover_image_hash, logo_hash, category, blockchain, created_by_id created_by, u.name creator_name").
			Joins("INNER JOIN users u ON u.user_id = created_by_id").
			Find(&domains).
			Error; err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Error("failed to get domains", err)
			return
		}
	}

	if len(domains) == 0 {
		httpo.NewErrorResponse(404, "no domains found").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "domains fetched successfully", domains).SendD(c)
}
