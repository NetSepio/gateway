package claim

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/claim")
	{
		g.POST("start", startClaimDomain)
		g.POST("finish", finishClaimDomain)
	}
}
