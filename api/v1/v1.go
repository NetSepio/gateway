package apiv1

import (
	authenticate "github.com/TheLazarusNetwork/netsepio-engine/api/v1/authenticate"
	claimrole "github.com/TheLazarusNetwork/netsepio-engine/api/v1/claimRole"
	delegatereviewcreation "github.com/TheLazarusNetwork/netsepio-engine/api/v1/delegateReviewCreation"
	flowid "github.com/TheLazarusNetwork/netsepio-engine/api/v1/flowid"
	"github.com/TheLazarusNetwork/netsepio-engine/api/v1/profile"
	roleid "github.com/TheLazarusNetwork/netsepio-engine/api/v1/roleId"

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
	}
}
