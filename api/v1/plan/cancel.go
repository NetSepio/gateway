package plan

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
)

func CancelStripeSubscription(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	db := dbconfig.GetDb()

	var user models.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "User not found").SendD(c)
		return
	}

	if user.SubscriptionStatus == "basic" || user.StripeSubscriptionId == nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "No active subscription to cancel").SendD(c)
		return
	}

	// Proceed to cancel the subscription
	_, err := subscription.Update(*user.StripeSubscriptionId, &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	})
	if err != nil {
		logwrapper.Errorf("Stripe subscription cancellation failed: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to cancel subscription").SendD(c)
		return
	}

	httpo.NewSuccessResponse(http.StatusOK, "Subscription cancelled successfully").SendD(c)
}
