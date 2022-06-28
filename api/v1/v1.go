package apiv1

import (
	authenticate "github.com/NetSepio/gateway/api/v1/authenticate"
	claimrole "github.com/NetSepio/gateway/api/v1/claimRole"
	delegatereviewcreation "github.com/NetSepio/gateway/api/v1/delegateReviewCreation"
	"github.com/NetSepio/gateway/api/v1/feedback"
	flowid "github.com/NetSepio/gateway/api/v1/flowid"
	"github.com/NetSepio/gateway/api/v1/healthcheck"
	"github.com/NetSepio/gateway/api/v1/profile"
	roleid "github.com/NetSepio/gateway/api/v1/roleId"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)
		profile.ApplyRoutes(v1)
		roleid.ApplyRoutes(v1)
		claimrole.ApplyRoutes(v1)
		delegatereviewcreation.ApplyRoutes(v1)
		healthcheck.ApplyRoutes(v1)
		feedback.ApplyRoutes(v1)
	}
}
