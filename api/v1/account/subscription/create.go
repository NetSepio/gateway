package subscription

import (
	"math"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/aptos"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"gorm.io/gorm"
)

func Buy111NFT(c *gin.Context) {
	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	if walletAddress == "" {
		logwrapper.Errorf("user has no wallet address")
		httpo.NewErrorResponse(http.StatusBadRequest, "user doesn't have any wallet linked").SendD(c)
		return
	}

	coinPrice, err := aptos.GetCoinPrice()
	if err != nil {
		logwrapper.Errorf("failed to get coin price: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(math.Ceil(coinPrice * 11.1 * 100))),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		logwrapper.Errorf("failed to create session: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	// insert in above table
	err = db.Create(&models.UserStripePi{
		Id:           uuid.NewString(),
		UserId:       userId,
		StripePiId:   pi.ID,
		StripePiType: models.Erebrus111NFT,
	}).Error
	if err != nil {
		logwrapper.Errorf("failed to insert into users_stripe_pi: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(http.StatusOK, "payment intent created", BuyErebrusNFTResponse{ClientSecret: pi.ClientSecret}).SendD(c)
}

func TrialSubscription(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)

	// Check if there is already an active trial subscription for the user
	var existingSubscription models.Subscription
	db := dbconfig.GetDb()
	if err := db.Where("user_id = ? AND type = ? AND end_time > ?", userId, "trial", time.Now()).First(&existingSubscription).Error; err == nil {
		// There is already an active trial subscription for the user
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already have an active trial subscription"})
		return
	}

	// Create a new trial subscription
	subscription := models.Subscription{
		UserId:    userId,
		StartTime: time.Now(),
		EndTime:   time.Now().AddDate(0, 0, 7),
		Type:      "TrialSubscription",
	}

	// Save the new trial subscription to the database
	if err := db.Model(models.Subscription{}).Create(&subscription).Error; err != nil {
		logwrapper.Errorf("Error creating subscription: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "subscription created"})
}

func CheckSubscription(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)

	db := dbconfig.GetDb()
	var subscription *models.Subscription
	err := db.Where("user_id = ?", userId).Order("end_time DESC").First(&subscription).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res := SubscriptionResponse{
				Status: "notFound",
			}
			c.JSON(http.StatusNotFound, res)
		}
		logwrapper.Errorf("Error fetching subscriptions: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	var status = "expired"
	if time.Now().Before(subscription.EndTime) {
		status = "active"
	}
	res := SubscriptionResponse{
		Subscription: subscription,
		Status:       status,
	}
	c.JSON(http.StatusOK, res)
}

func CreatePaymentIntent(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	db := dbconfig.GetDb()
	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(1000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Payment to purchase 3 month subscription"),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		logwrapper.Errorf("failed to create new payment intent: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	// insert in above table
	err = db.Create(&models.UserStripePi{
		Id:           uuid.NewString(),
		UserId:       userId,
		StripePiId:   pi.ID,
		StripePiType: models.ThreeMonthSubscription,
	}).Error
	if err != nil {
		logwrapper.Errorf("failed to insert into users_stripe_pi: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Created new charge", gin.H{"clientSecret": pi.ClientSecret}).SendD(c)
}
