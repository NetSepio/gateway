package domain

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func postDomain(c *gin.Context) {
	db := dbconfig.GetDb()
	var request CreateDomainRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	domainId := uuid.NewString()
	txtValue := fmt.Sprintf("netsepio_verification=%s", uuid.NewString())
	newDomain := models.Domain{
		Id:               domainId,
		TxtValue:         &txtValue,
		DomainName:       request.DomainName,
		Title:            request.Title,
		Headline:         request.Headline,
		Description:      request.Description,
		LogoHash:         request.LogoHash,
		Category:         request.Category,
		CoverImageHash:   request.CoverImageHash,
		Blockchain:       request.Blockchain,
		CreatedByAddress: strings.ToLower(walletAddress),
		UpdatedByAddress: strings.ToLower(walletAddress),
	}

	domainAdmin := models.DomainAdmin{
		DomainId:           domainId,
		AdminWalletAddress: strings.ToLower(walletAddress),
		UpdatedByAddress:   strings.ToLower(walletAddress),
		Name:               request.AdminName,
		Role:               request.AdminRole,
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := db.Create(&newDomain).Error; err != nil {
			logwrapper.Errorf("failed to create domain: %s", err)
			return err
		}
		if err := db.Create(&domainAdmin).Error; err != nil {
			logwrapper.Errorf("failed to associate admin with domain: %s", err)
			return err
		}

		return nil
	})

	if err != nil {
		logwrapper.Errorf("failed to create domain: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create domain").SendD(c)
		return
	}
	payload := CreateDomainResponse{
		TxtValue: txtValue, DomainId: domainId,
	}
	httpo.NewSuccessResponseP(200, "domain created", payload).SendD(c)
}