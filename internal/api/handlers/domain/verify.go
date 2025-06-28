package domain

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
)

func verifyDomain(c *gin.Context) {
	db := database.GetDb()
	var request VerifyDomainRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

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

	if *domainData.TxtValue == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "request verification before verifying").SendD(c)
		return
	}
	//dig +short gmail.com txt
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

	ttrue := true
	strempty := ""
	for _, txt := range txts {
		if txt == *domainData.TxtValue {
			domainData.Verified = &ttrue
			domainData.TxtValue = &strempty
			if err := db.Model(&models.Domain{}).
				Where("domain_name = ?", domainData.DomainName).
				Update("verified", false).
				Error; err != nil {
				logwrapper.Errorf("failed to mark other domain records as unverified: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "failed to verify Domain").SendD(c)
				return
			}
			if err := db.Updates(&domainData).Error; err != nil {
				logwrapper.Errorf("failed to mark domain as verified: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "failed to verify Domain").SendD(c)
				return
			}
			httpo.NewSuccessResponse(200, "domain verified").SendD(c)
			return
		}
	}
	httpo.NewErrorResponse(400, "no valid txt record found").SendD(c)
}
