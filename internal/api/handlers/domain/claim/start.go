package claim

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func startClaimDomain(c *gin.Context) {
	db := database.GetDb()
	var request ClaimDomainRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	var domain models.Domain
	err = db.Where("id = ?", request.DomainId).First(&domain).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusNotFound, "domain not found").SendD(c)
		return
	}

	if !domain.Claimable {
		httpo.NewErrorResponse(http.StatusForbidden, "domain is not claimable").SendD(c)
		return
	}

	txtValue := fmt.Sprintf("netsepio_verification=%s", uuid.NewString())
	userId := c.GetString(paseto.CTX_USER_ID)
	err = db.Transaction(func(tx *gorm.DB) error {
		// insert txt record in new model
		newClaimTxt := models.DomainClaim{
			ID:       uuid.NewString(),
			DomainID: request.DomainId,
			Txt:      txtValue,
			AdminId:  userId,
		}

		//remove all txt claim records for this domain and current user as admin
		if err := tx.Where("domain_id = ? and admin_id = ?", request.DomainId, userId).Delete(&models.DomainClaim{}).Error; err != nil {
			return fmt.Errorf("failed to delete claim txt: %s", err)
		}
		// insert into db
		if err := tx.Create(&newClaimTxt).Error; err != nil {
			return fmt.Errorf("failed to create claim txt: %s", err)
		}

		return nil
	})

	if err != nil {
		logwrapper.Errorf("failed to create claim txt: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	//return txt record
	payload := ClaimDomainResponse{
		TxtValue: txtValue,
	}
	httpo.NewSuccessResponseP(200, "domain claimed", payload).SendD(c)

}
