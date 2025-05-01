package domain

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func patchDomain(c *gin.Context) {
	db := database.GetDb()
	var requestBody PatchDomainRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}
	userId := c.GetString(paseto.CTX_USER_ID) // corrected 'paset' to 'paseto'
	err = db.Model(&models.DomainAdmin{}).
		Where(&models.DomainAdmin{DomainId: requestBody.DomainId, AdminId: userId}).
		First(&models.DomainAdmin{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not exist or user is not admin of the domain").SendD(c)
			return
		}

		logwrapper.Errorf("failed to get domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get admins").SendD(c)
	}
	domainUpdate := models.Domain{
		Title:          requestBody.Title,
		Headline:       requestBody.Headline,
		Description:    requestBody.Description,
		LogoHash:       requestBody.LogoHash,
		Category:       requestBody.Category,
		CoverImageHash: requestBody.CoverImageHash,
		Blockchain:     requestBody.Blockchain,
		UpdatedById:    userId,
	}
	result := db.Model(&models.Domain{}).
		Where("id = ?", requestBody.DomainId).
		Updates(&domainUpdate)
	if result.Error != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "unexpected error occured").SendD(c)
		return
	}
	if result.RowsAffected == 0 {
		httpo.NewErrorResponse(http.StatusNotFound, "domain not found").SendD(c)
		return
	}
	httpo.NewSuccessResponse(200, "domain successfully updated").SendD(c)

}
