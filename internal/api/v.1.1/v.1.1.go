package v11

import (
	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/handlers/account"
	"netsepio-gateway-v1.1/internal/api/handlers/agents"
	"netsepio-gateway-v1.1/internal/api/handlers/agents/cyreneAiAgent"
	"netsepio-gateway-v1.1/internal/api/handlers/authenticate"
	"netsepio-gateway-v1.1/internal/api/handlers/certificate"
	"netsepio-gateway-v1.1/internal/api/handlers/client"
	delegatereviewcreation "netsepio-gateway-v1.1/internal/api/handlers/delegateReviewCreation"
	"netsepio-gateway-v1.1/internal/api/handlers/deletereview"
	"netsepio-gateway-v1.1/internal/api/handlers/domain"
	"netsepio-gateway-v1.1/internal/api/handlers/feedback"
	"netsepio-gateway-v1.1/internal/api/handlers/flowid"
	"netsepio-gateway-v1.1/internal/api/handlers/getreviewerdetails"
	"netsepio-gateway-v1.1/internal/api/handlers/getreviews"
	"netsepio-gateway-v1.1/internal/api/handlers/leaderboard"
	"netsepio-gateway-v1.1/internal/api/handlers/nftcontract"
	nodedwifi "netsepio-gateway-v1.1/internal/api/handlers/nodeDwifi"
	"netsepio-gateway-v1.1/internal/api/handlers/nodes"
	"netsepio-gateway-v1.1/internal/api/handlers/organisation"
	"netsepio-gateway-v1.1/internal/api/handlers/perks"
	"netsepio-gateway-v1.1/internal/api/handlers/profile"
	"netsepio-gateway-v1.1/internal/api/handlers/referral"
	"netsepio-gateway-v1.1/internal/api/handlers/registerDwifi"
	"netsepio-gateway-v1.1/internal/api/handlers/report"
	"netsepio-gateway-v1.1/internal/api/handlers/sdkauthentication"
	caddyservices "netsepio-gateway-v1.1/internal/api/handlers/services"
	"netsepio-gateway-v1.1/internal/api/handlers/stats"
	"netsepio-gateway-v1.1/internal/api/handlers/status"
	"netsepio-gateway-v1.1/internal/api/handlers/subscription"
	"netsepio-gateway-v1.1/internal/api/handlers/summary"
	"netsepio-gateway-v1.1/internal/api/handlers/waitlist"
	"netsepio-gateway-v1.1/internal/api/handlers/walrus"
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
	v1 := r.Group("/v1.1")
	{
		organisation.ApplyRoutes(v1)
		certificate.ApplyRoutes(v1)
		subscription.ApplyRoutesV11(v1)
	}
}
