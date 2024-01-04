package plan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

func updateSubscriptionStatus(customerID, subscriptionStatus string, stripeSubscriptionId *string, stripeSubscriptionStatus stripe.SubscriptionStatus) error {
	db := dbconfig.GetDb()
	var user models.User
	if err := db.Where("stripe_customer_id = ?", customerID).First(&user).Error; err != nil {
		return err
	}

	user.StripeSubscriptionId = stripeSubscriptionId
	user.SubscriptionStatus = subscriptionStatus
	user.StripeSubscriptionStatus = stripeSubscriptionStatus
	return db.Save(&user).Error
}

func StripeWebhookHandler(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logwrapper.Errorf("Error reading request body: %v", err)
		c.Status(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), envconfig.EnvVars.STRIPE_WEBHOOK_SECRET)
	if err != nil {
		logwrapper.Errorf("Error verifying webhook signature: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}
	switch event.Type {
	case stripe.EventTypeCustomerSubscriptionDeleted:
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if err := updateSubscriptionStatus(subscription.Customer.ID, "basic", nil, "unset"); err != nil {
			logwrapper.Errorf("Error updating subscription status: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

	case stripe.EventTypeCustomerSubscriptionUpdated:
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			// w.WriteHeader(http.StatusBadRequest)
			return
		}
		if subscription.Status == "active" {
			if err := updateSubscriptionStatus(subscription.Customer.ID, subscription.Items.Data[0].Price.LookupKey, &subscription.ID, subscription.Status); err != nil {
				logwrapper.Errorf("Error updating subscription status: %v", err)
				c.Status(http.StatusInternalServerError)
				return
			}
		} else {
			if err := updateSubscriptionStatus(subscription.Customer.ID, "basic", &subscription.ID, subscription.Status); err != nil {
				logwrapper.Errorf("Error updating subscription status: %v", err)
				c.Status(http.StatusInternalServerError)
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "received"})
}
