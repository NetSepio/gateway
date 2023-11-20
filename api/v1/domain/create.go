package domain

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	domainId := uuid.NewString()
	txtValue := fmt.Sprintf("netsepio_verification=%s", uuid.NewString())
	newDomain := &models.Domain{
		Id:             domainId,
		TxtValue:       &txtValue,
		DomainName:     request.DomainName,
		Title:          request.Title,
		Headline:       request.Headline,
		Description:    request.Description,
		LogoHash:       request.LogoHash,
		Category:       request.Category,
		CoverImageHash: request.CoverImageHash,
	}

	payload := CreateDomainResponse{
		TxtValue: txtValue, DomainId: domainId,
	}

	if err := db.Create(newDomain).Error; err != nil {
		logwrapper.Errorf("failed to create domain: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create domain").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "domain created", payload).SendD(c)
}
