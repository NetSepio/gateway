package report

import (
	"net/http"
	"time"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// getReports fetches reports with optional filters
func getReports(c *gin.Context) {
	var filter ReportFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "Invalid query parameters").SendD(c)
		return
	}

	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)

	// Query with vote counts
	query := db.Debug().Model(&models.Report{}).
		Select(`reports.*, 
			(SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and report_votes.vote_type = 'upvote') as upvotes,
			(SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and report_votes.vote_type = 'downvote') as downvotes,
			(SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and report_votes.vote_type = 'notsure') as notSure,
			(SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id) as totalvotes,
			reports.end_time,
			(SELECT vote_type FROM report_votes WHERE report_id = reports.id AND voter_id = ?) as user_vote`, userId).
		Joins("LEFT JOIN report_votes ON report_votes.report_id = reports.id").
		Group("reports.id")

	// Apply filters
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Title+"%")
	}
	if filter.ProjectDomain != "" {
		query = query.Where("project_domain = ?", filter.ProjectDomain)
	}
	if filter.ProjectName != "" {
		query = query.Where("project_name = ?", filter.ProjectName)
	}

	// Execute query
	var reportsWithVotes []ReportWithVotes
	if err := query.Find(&reportsWithVotes).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch reports").SendD(c)
		return
	}

	// Calculate status for each report
	for i := range reportsWithVotes {
		report := &reportsWithVotes[i]
		if time.Now().After(report.EndTime) {
			if float64(report.Upvotes)/float64(report.Totalvotes) >= 0.51 {
				report.Status = "accepted"
			} else {
				report.Status = "rejected"
			}
		} else {
			report.Status = "running"
		}
	}

	httpo.NewSuccessResponseP(http.StatusOK, "Reports fetched successfully", reportsWithVotes).SendD(c)
}
