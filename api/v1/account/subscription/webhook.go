package subscription

import (
	"encoding/json"
	"errors"
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
	"gorm.io/gorm"
)

func StripeWebhookHandler(c *gin.Context) {
	db := dbconfig.GetDb()

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
	case stripe.EventTypePaymentIntentSucceeded:
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// get user with stripe_pi_id
		var user models.User
		if err := db.Where("stripe_pi_id = ?", paymentIntent.ID).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//warn and return success
				logwrapper.Warnf("No user found with stripe_pi_id: %v", err)
				c.JSON(http.StatusOK, gin.H{"status": "received"})
				return
			}
			logwrapper.Errorf("Error getting user with stripe_pi_id: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if err := db.Model(&models.User{}).Where("stripe_pi_id = ?", paymentIntent.ID).Update("stripe_session_id", nil).Error; err != nil {
			logwrapper.Errorf("Error updating stripe_pi_id: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("minting nft -- 111NFT")

	case stripe.EventTypePaymentIntentCanceled:
		err := HandleCanceledOrFailedPaymentIntent(event.Data.Raw)
		if err != nil {
			logwrapper.Errorf("Error handling canceled payment intent: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "received"})
}

func HandleCanceledOrFailedPaymentIntent(eventDataRaw json.RawMessage) error {
	db := dbconfig.GetDb()
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(eventDataRaw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error parsing webhook JSON: %w", err)
	}

	// TODO rows affected
	if err := db.Model(&models.User{}).Where("stripe_pi_id = ?", paymentIntent.ID).Update("stripe_session_id", nil).Error; err != nil {
		return fmt.Errorf("error updating stripe_pi_id: %w", err)
	}

	return nil
}
