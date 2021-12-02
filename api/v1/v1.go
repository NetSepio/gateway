package apiv1

import (
	authenticate "netsepio-api/api/v1/authenticate"
	flowid "netsepio-api/api/v1/flowid"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)
	}
}
