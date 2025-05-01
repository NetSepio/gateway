package domain

import (
	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/handlers/domain/admin"
	"netsepio-gateway-v1.1/internal/api/handlers/domain/claim"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
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
