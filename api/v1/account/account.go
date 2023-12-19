package account

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/account")
	{
		g.POST("auth-google", authGoogle)
	}
}
