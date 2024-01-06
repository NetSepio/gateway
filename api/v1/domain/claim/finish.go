package claim

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// finish claim, check if txt record in dns matches the domain name, if matches then add current user to admin
// verify txt using dig +short gmail.com txt check above code for refereance
func finishClaimDomain(c *gin.Context) {
	db := dbconfig.GetDb()
	var request FinishClaimDomainRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	// fetch domainData from db
	var domainData models.Domain
	err = db.Model(&domainData).Where("id = ?", request.DomainId).First(&domainData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not found").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	userId := c.GetString(paseto.CTX_USER_ID)
	var claimTxt models.DomainClaim
	err = db.Where("domain_id = ? and admin_id = ?", request.DomainId, userId).First(&claimTxt).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, "claim record not found").SendD(c)
		return
	}

	var domainAdmin models.DomainAdmin
	err = db.Where("domain_id = ?", request.DomainId).First(&domainAdmin).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusForbidden, "domain is already claimed").SendD(c)
		return
	}
	if err != gorm.ErrRecordNotFound {
		logwrapper.Errorf("failed to query domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusNotFound, "domain admin not found").SendD(c)
		return
	}

	txts, err := net.LookupTXT(domainData.DomainName)
	if err != nil {
		var dnsError *net.DNSError
		if errors.As(err, &dnsError) {
			if dnsError.IsNotFound {
				httpo.NewErrorResponse(http.StatusNotFound, "record not found in dns").SendD(c)
				return
			}
		}
		logwrapper.Errorf("failed to lookup domain: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	if len(txts) == 0 {
		httpo.NewErrorResponse(http.StatusNotFound, "txt records are empty").SendD(c)
		return
	}

	// Add current user to admin
	newDomainAdmin := models.DomainAdmin{
		DomainId:    request.DomainId,
		AdminId:     userId,
		UpdatedById: userId,
	}

	// Insert into db
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newDomainAdmin).Error; err != nil {
			return fmt.Errorf("failed to create domain admin: %s", err)
		}

		// update domain set verified to true
		if err := tx.Model(&domainData).Where("id = ?", request.DomainId).Update("verified", true).Update("claimable", false).Error; err != nil {
			return fmt.Errorf("failed to update domain: %s", err)
		}

		return nil
	})

	if err != nil {
		logwrapper.Errorf("failed to complete transaction: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "domain claim completed").SendD(c)
}
