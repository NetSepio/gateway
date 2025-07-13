package v11

import (
	"github.com/NetSepio/gateway/internal/api/handlers/account"
	"github.com/NetSepio/gateway/internal/api/handlers/agents"
	"github.com/NetSepio/gateway/internal/api/handlers/agents/cyreneAiAgent"
	"github.com/NetSepio/gateway/internal/api/handlers/authenticate"
	"github.com/NetSepio/gateway/internal/api/handlers/certificate"
	"github.com/NetSepio/gateway/internal/api/handlers/client"
	delegatereviewcreation "github.com/NetSepio/gateway/internal/api/handlers/delegateReviewCreation"
	"github.com/NetSepio/gateway/internal/api/handlers/deletereview"
	"github.com/NetSepio/gateway/internal/api/handlers/domain"
	"github.com/NetSepio/gateway/internal/api/handlers/feedback"
	"github.com/NetSepio/gateway/internal/api/handlers/flowid"
	"github.com/NetSepio/gateway/internal/api/handlers/getreviewerdetails"
	"github.com/NetSepio/gateway/internal/api/handlers/getreviews"
	"github.com/NetSepio/gateway/internal/api/handlers/leaderboard"
	"github.com/NetSepio/gateway/internal/api/handlers/nftcontract"
	nodedwifi "github.com/NetSepio/gateway/internal/api/handlers/nodeDwifi"
	"github.com/NetSepio/gateway/internal/api/handlers/nodes"
	"github.com/NetSepio/gateway/internal/api/handlers/organisation"
	"github.com/NetSepio/gateway/internal/api/handlers/perks"
	"github.com/NetSepio/gateway/internal/api/handlers/profile"
	"github.com/NetSepio/gateway/internal/api/handlers/referral"
	"github.com/NetSepio/gateway/internal/api/handlers/registerDwifi"
	"github.com/NetSepio/gateway/internal/api/handlers/report"
	"github.com/NetSepio/gateway/internal/api/handlers/sdkauthentication"
	caddyservices "github.com/NetSepio/gateway/internal/api/handlers/services"
	"github.com/NetSepio/gateway/internal/api/handlers/stats"
	"github.com/NetSepio/gateway/internal/api/handlers/status"
	"github.com/NetSepio/gateway/internal/api/handlers/subscription"
	"github.com/NetSepio/gateway/internal/api/handlers/summary"
	"github.com/NetSepio/gateway/internal/api/handlers/waitlist"
	"github.com/NetSepio/gateway/internal/api/handlers/walrus"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)

		authenticate.ApplyRoutes(v1)
		profile.ApplyRoutes(v1)
		delegatereviewcreation.ApplyRoutes(v1)
		deletereview.ApplyRoutes(v1)
		feedback.ApplyRoutes(v1)
		waitlist.ApplyRoutes(v1)
		stats.ApplyRoutes(v1)
		getreviews.ApplyRoutes(v1)
		getreviewerdetails.ApplyRoutes(v1)
		domain.ApplyRoutes(v1)
		report.ApplyRoutes(v1)
		account.ApplyRoutes(v1)
		// siteinsights.ApplyRoutes(v1)
		summary.ApplyRoutes(v1)
		sdkauthentication.ApplyRoutes(v1)
		leaderboard.ApplyRoutes(v1)
		nftcontract.ApplyRoutes(v1)
		referral.ApplyReferraAccountlRoutes(v1)

		// erebrus
		status.ApplyRoutes(v1)
		client.ApplyRoutes(v1)
		nodes.ApplyRoutes(v1)
		subscription.ApplyRoutes(v1)
		registerDwifi.ApplyRoutes(v1)
		nodedwifi.ApplyRoutes(v1)
		walrus.ApplyRoutes(v1)
		caddyservices.ApplyRoutes(v1)
		agents.ApplyRoutes(v1)
		cyreneAiAgent.ApplyRoutes(v1)
		perks.ApplyRoutesPerks(v1)

	}
}

func ApplyRoutesV1_1(r *gin.RouterGroup) {
	v11 := r.Group("/v1.1")
	{
		organisation.ApplyRoutes(v11)
		certificate.ApplyRoutes(v11)
		subscription.ApplyRoutesV11(v11)
		profile.ApplyRoutesv11(v11)
		agents.ApplyRoutesv11(v11)
	}
}
