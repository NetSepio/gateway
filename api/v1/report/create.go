package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/aptos"
	"github.com/NetSepio/gateway/util/pkg/ipfs"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExtendedReport struct {
	models.Report
	Tags   []string `json:"tags"`
	Images []string `json:"images"`
}

func postReport(c *gin.Context) {
	var request ReportRequest
	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Invalid request body: %s", err)).SendD(c)
		return
	}

	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)              // Get user ID from context
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES) // Get user ID from context
	newReport := models.Report{
		ID:            uuid.NewString(),
		Title:         request.Title,
		Description:   request.Description,
		Document:      request.Document,
		ProjectName:   request.ProjectName,
		ProjectDomain: request.ProjectDomain,
		CreatedBy:     userId,
		Category:      request.Category,
		EndTime:       time.Now().Add(time.Hour * 24 * 2),
		Status:        "running",
	}
	extendedReport := struct {
		models.Report
		Tags   []string `json:"tags"`
		Images []string `json:"images"`
	}{
		Report: newReport,
		Tags:   request.Tags,
		Images: request.Images,
	}
	// Convert report to JSON for IPFS upload
	reportJSON, err := json.Marshal(extendedReport)
	if err != nil {
		logwrapper.Errorf("failed to marshal report: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	// Upload to IPFS
	ipfsResponse, err := ipfs.UploadToIpfs(bytes.NewReader(reportJSON), "report.json")
	if err != nil {
		logwrapper.Errorf("failed to upload to IPFS: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	// Store the IPFS hash in the report object
	newReport.MetaDataHash = &ipfsResponse.Value.Cid

	txResult, err := aptos.SubmitProposal(walletAddress, *newReport.MetaDataHash)
	if err != nil {
		logwrapper.Errorf("failed to submit proposal: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	newReport.TransactionHash = &txResult.Result.TransactionHash
	newReport.TransactionVersion = &txResult.Result.Version
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newReport).Error; err != nil {
			return fmt.Errorf("failed to insert report: %w", err)
		}

		// Insert tags
		for _, tag := range request.Tags {
			if err := tx.Create(&models.ReportTag{ReportID: newReport.ID, Tag: tag}).Error; err != nil {
				return fmt.Errorf("failed to insert tag for report: %w", err)
			}
		}

		// Insert images
		for _, imageURL := range request.Images {
			if err := tx.Create(&models.ReportImage{ReportID: newReport.ID, ImageURL: imageURL}).Error; err != nil {
				return fmt.Errorf("failed to insert image for report: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		logwrapper.Errorf("failed to create report: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create report").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Report created successfully", newReport).SendD(c)
}
