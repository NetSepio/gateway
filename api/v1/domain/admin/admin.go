package admin

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/admin")
	{
		g.GET("", getAdmin)
		g.POST("", createAdmin)
		g.DELETE("", deleteAdmin)
	}
}
