package plan

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	plan := r.Group("/plan")
	{
		plan.POST("/webhook", StripeWebhookHandler)
		plan.Use(paseto.PASETO(false))
		plan.POST("/", CreateStripeSession)
		plan.DELETE("/", CancelStripeSubscription)
	}
}
