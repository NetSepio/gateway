package report

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func postReport(c *gin.Context) {
	var request ReportRequest
	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Invalid request body: %s", err)).SendD(c)
		return
	}

	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID) // Get user ID from context
	newReport := models.Report{
		ID:            uuid.NewString(),
		Title:         request.Title,
		Description:   request.Description,
		Document:      request.Document,
		ProjectName:   request.ProjectName,
		ProjectDomain: request.ProjectDomain,
		CreatedBy:     userId,
		EndTime:       time.Now().Add(time.Hour * 24 * 2),
	}
	err := db.Transaction(func(tx *gorm.DB) error {
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
