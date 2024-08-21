package apiv1

import (
	"github.com/NetSepio/gateway/api/v1/account"
	"github.com/NetSepio/gateway/api/v1/account/subscription"
	authenticate "github.com/NetSepio/gateway/api/v1/authenticate"
	delegatereviewcreation "github.com/NetSepio/gateway/api/v1/delegateReviewCreation"
	"github.com/NetSepio/gateway/api/v1/deletereview"
	"github.com/NetSepio/gateway/api/v1/domain"
	"github.com/NetSepio/gateway/api/v1/dvpnnft"
	"github.com/NetSepio/gateway/api/v1/erebrus"
	"github.com/NetSepio/gateway/api/v1/feedback"
	flowid "github.com/NetSepio/gateway/api/v1/flowid"
	"github.com/NetSepio/gateway/api/v1/getreviewerdetails"
	"github.com/NetSepio/gateway/api/v1/getreviews"
	"github.com/NetSepio/gateway/api/v1/leaderboard"
	"github.com/NetSepio/gateway/api/v1/nftcontract"
	"github.com/NetSepio/gateway/api/v1/profile"
	"github.com/NetSepio/gateway/api/v1/report"
	"github.com/NetSepio/gateway/api/v1/sdkauthentication"
	"github.com/NetSepio/gateway/api/v1/siteinsights"
	"github.com/NetSepio/gateway/api/v1/sotreus"
	"github.com/NetSepio/gateway/api/v1/stats"
	"github.com/NetSepio/gateway/api/v1/status"
	"github.com/NetSepio/gateway/api/v1/summary"
	"github.com/NetSepio/gateway/api/v1/waitlist"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)
		profile.ApplyRoutes(v1)
		delegatereviewcreation.ApplyRoutes(v1)
		deletereview.ApplyRoutes(v1)
		status.ApplyRoutes(v1)
		feedback.ApplyRoutes(v1)
		waitlist.ApplyRoutes(v1)
		stats.ApplyRoutes(v1)
		getreviews.ApplyRoutes(v1)
		getreviewerdetails.ApplyRoutes(v1)
		sotreus.ApplyRoutes(v1)
		domain.ApplyRoutes(v1)
		erebrus.ApplyRoutes(v1)
		report.ApplyRoutes(v1)
		account.ApplyRoutes(v1)
		siteinsights.ApplyRoutes(v1)
		subscription.ApplyRoutes(v1)
		summary.ApplyRoutes(v1)
		sdkauthentication.ApplyRoutes(v1)
		leaderboard.ApplyRoutes(v1)
		nftcontract.ApplyRoutes(v1)
		dvpnnft.ApplyRoutes(v1)

	}
}
