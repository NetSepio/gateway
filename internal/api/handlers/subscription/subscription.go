package subscription

import (
	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.POST("payment-webhook", StripeWebhookHandler)
		g.Use(paseto.PASETO(false))
		g.POST("/trial", TrialSubscription)
		g.PATCH("/trial", PatchTrialSubscription)
		g.POST("/create-payment", CreatePaymentIntent)
		g.GET("", CheckSubscription)
		g.POST("erebrus", Buy111NFT)
	}
}
