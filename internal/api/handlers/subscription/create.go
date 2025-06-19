package subscription

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
	"netsepio-gateway-v1.1/utils/pkg/aptos"
)

func Buy111NFT(c *gin.Context) {
	db := database.GetDB2()
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

	if len(userId) == 0 {
		userId = c.GetString(paseto.CTX_ORGANISATION_ID)
	}

	// Check if there is already an active trial subscription for the user
	var existingSubscription models.Subscription
	db := database.GetDB2()
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

func PatchTrialSubscription(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)

	// Check if there is already an active trial subscription for the user
	var existingSubscription models.Subscription
	db := database.GetDB2()
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

	ordId := c.GetString(paseto.CTX_ORGANISATION_ID)

	if len(userId) == 0 {
		userId = ordId
	}

	db := database.GetDB2()
	var subscription *models.Subscription
	err := db.Where("user_id = ?", userId).Order("end_time DESC").First(&subscription).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res := SubscriptionResponse{
				Status: "subscription not found",
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
	db := database.GetDB2()
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
