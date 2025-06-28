package subscription

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/NetSepio/gateway/utils/pkg/aptos"
)

func StripeWebhookHandler(c *gin.Context) {
	db := database.GetDb()

	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logwrapper.Errorf("Error reading request body: %v", err)
		c.Status(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), load.Cfg.STRIPE_WEBHOOK_SECRET)
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

		// Get user with stripe_pi_id
		var userStripePi models.UserStripePi
		if err := db.Where("stripe_pi_id = ?", paymentIntent.ID).First(&userStripePi).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Warn and return success
				logwrapper.Warnf("No user found with stripe_pi_id: %v", err)
				c.JSON(http.StatusOK, gin.H{"status": "received"})
				return
			}
			logwrapper.Errorf("Error getting user with stripe_pi_id: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// Get user with user_id
		var user models.User
		if err := db.Where("user_id = ?", userStripePi.UserId).First(&user).Error; err != nil {
			logwrapper.Errorf("Error getting user with user_id: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		if userStripePi.StripePiType == models.ThreeMonthSubscription {
			// Handle ThreeMonthSubscription payment intent
			subscription := models.Subscription{
				UserId:    user.UserId,
				StartTime: time.Now(),
				EndTime:   time.Now().AddDate(0, 3, 0),
				Type:      "ThreeMonthSubscription",
			}
			if err = db.Model(models.Subscription{}).Create(&subscription).Error; err != nil {
				logwrapper.Errorf("Error creating subscription: %v", err)
				c.Status(http.StatusInternalServerError)
				return
			}
		} else if userStripePi.StripePiType == models.Erebrus111NFT {
			// Handle Erebrus111NFT payment intent
			if _, err = aptos.DelegateMintNft(*user.WalletAddress); err != nil {
				logwrapper.Errorf("Error minting nft: %v", err)
				c.Status(http.StatusInternalServerError)
				return
			}
			fmt.Println("minting nft -- 111NFT")
		}

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
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(eventDataRaw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error parsing webhook JSON: %w", err)
	}

	return nil
}
