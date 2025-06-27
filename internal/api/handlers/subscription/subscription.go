package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
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
		g.GET("/list", getAllSubscription)
	}
}
func ApplyRoutesV11(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.Use(paseto.PASETO(false))
		g.POST("", CreateSubscription)
	}
}
