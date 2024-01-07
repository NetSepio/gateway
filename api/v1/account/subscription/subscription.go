package subscription

import (
	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.POST("stripe-webhook", StripeWebhookHandler)
		g.Use(paseto.PASETO(false))
		g.POST("111-nft", Buy111NFT)
	}
}
