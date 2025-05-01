package report

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

type ReportVoteRequest struct {
	ReportID string `json:"reportId" binding:"required"`
	VoteType string `json:"voteType" binding:"required,oneof=upvote downvote notsure"`
}

func postReportVote(c *gin.Context) {
	var request ReportVoteRequest
	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Invalid request body: %s", err)).SendD(c)
		return
	}

	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID) // Assuming user ID is in the context

	// Check if the voting period has ended
	var report models.Report
	if err := db.Where("id = ?", request.ReportID).First(&report).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to vote").SendD(c)
		return
	}

	if time.Now().After(report.EndTime) {
		httpo.NewErrorResponse(http.StatusBadRequest, "Voting period has ended").SendD(c)
		return
	}

	// Insert or update the vote
	newVote := models.ReportVote{
		ReportID: request.ReportID,
		VoterID:  userId,
		VoteType: request.VoteType,
	}
	err := db.Create(&newVote).Error
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == "23505" && pgError.ConstraintName == "report_votes_pkey" {
				httpo.NewErrorResponse(http.StatusBadRequest, "You have already voted on this report").SendD(c)
				return
			}
		}
		logwrapper.Errorf("failed to record vote: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to record vote").SendD(c)
		return
	}

	httpo.NewSuccessResponse(http.StatusOK, "Vote recorded successfully").SendD(c)
}
