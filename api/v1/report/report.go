package report

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	report := r.Group("/report")
	{
		report.Use(paseto.PASETO(false))

		report.POST("/", postReport)         // Endpoint for creating a new report
		report.GET("/", getReports)          // Endpoint for retrieving reports with optional filters
		report.POST("/vote", postReportVote) // Endpoint for voting on a report
	}
}
