package report

import (
	"time"

	"github.com/google/uuid"
)

// ReportRequest defines the structure for report creation request
type ReportRequest struct {
	Title         string   `json:"title" binding:"required"`
	Description   string   `json:"description"`
	Images        []string `json:"image"`
	Document      string   `json:"document"`
	Category      string   `json:"category"`
	Tags          []string `json:"tags"`
	ProjectName   string   `json:"projectName"`
	ProjectDomain string   `json:"projectDomain"`
}

// ReportFilter for query parameters
type ReportFilter struct {
	Title         string `form:"title"`
	ProjectDomain string `form:"projectDomain"`
	ProjectName   string `form:"projectName"`
	Accepted      *bool  `form:"accepted"`
}

// ReportWithVotes and calculated status
type ReportWithVotes struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Document      string    `json:"document"`
	ProjectName   string    `json:"projectName"`
	ProjectDomain string    `json:"projectDomain"`
	CreatedBy     uuid.UUID `json:"createdBy"`
	EndTime       time.Time `json:"endTime"`
	Upvotes       int       `json:"upvotes"`
	Downvotes     int       `json:"downvotes"`
	Notsure       int       `json:"notsure"`
	Totalvotes    int       `json:"totalVotes"`
	Status        string    `json:"status"` // Calculated status
	UserVote      string    `json:"userVote"`
	Category      string    `json:"category"`
}
