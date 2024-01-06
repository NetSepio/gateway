package domain

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/api/v1/domain/admin"
	"github.com/NetSepio/gateway/api/v1/domain/claim"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/domain")
	{
		g.Use(paseto.PASETO(true))
		g.GET("", queryDomain)
		g.Use(paseto.PASETO(false))
		g.POST("", postDomain)
		g.DELETE("", deleteDomain)
		g.PATCH("", patchDomain)
		g.PATCH("/verify", verifyDomain)
		claim.ApplyRoutes(g)
		admin.ApplyRoutes(g)
	}
}
