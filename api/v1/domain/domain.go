package domain

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/domain")
	{
		g.Use(paseto.PASETO)
		g.POST("", postDomain)
		g.GET("", queryDomain)
		g.DELETE("", deleteDomain)
		g.PATCH("", patchDomain)
		g.PATCH("/verify", verifyDomain)
	}
}
