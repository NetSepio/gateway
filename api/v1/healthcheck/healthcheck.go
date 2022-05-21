package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/healthcheck")
	{
		g.GET("", healthCheck)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "alive",
	})
}
