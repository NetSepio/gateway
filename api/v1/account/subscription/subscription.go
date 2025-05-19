package subscription

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.POST("payment-webhook", StripeWebhookHandler)
		g.Use(paseto.PASETO(false))
		g.POST("/trial", TrialSubscription)
		g.PATCH("/trial", TrialSubscription)
		g.POST("/create-payment", CreatePaymentIntent)
		g.GET("", CheckSubscription)
		g.POST("erebrus", Buy111NFT)
		g.POST("/custom_duration", SubscriptionForCustomDuration)
	}
}
