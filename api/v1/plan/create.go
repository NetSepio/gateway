package plan

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
)

func CreateStripeSession(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	var req struct {
		PriceID string `json:"priceId" binding:"required"`
	}
	if err := c.BindJSON(&req); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err)).SendD(c)
		return
	}
	db := dbconfig.GetDb()
	var user models.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		logwrapper.Errorf("Failed to find user: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	var customerID string

	// Check if stripe_customer_id is null
	if user.StripeCustomerId == "" {
		// Create a new Stripe customer
		customerParams := &stripe.CustomerParams{}
		stripeCustomer, err := customer.New(customerParams)
		if err != nil {
			logwrapper.Errorf("Stripe customer creation failed: %v", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
		customerID = stripeCustomer.ID

		// Update user with new stripe_customer_id
		user.StripeCustomerId = stripeCustomer.ID
		if err := db.Save(&user).Error; err != nil {
			logwrapper.Errorf("Failed to update user: %v", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
	} else {
		customerID = user.StripeCustomerId
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(req.PriceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:              stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL:        stripe.String(envconfig.EnvVars.STRIPE_SUCCESS_URL),
		CancelURL:         stripe.String(envconfig.EnvVars.STRIPE_CANCEL_URL),
		ClientReferenceID: stripe.String(userId),
		Customer:          stripe.String(customerID),
	}

	s, err := session.New(params)
	if err != nil {
		logwrapper.Errorf("Stripe session creation failed: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create Stripe session").SendD(c)
		return
	}

	fmt.Println("customer ", s.Customer)

	httpo.NewSuccessResponseP(http.StatusOK, "Session created successfully", gin.H{"session_url": s.URL}).SendD(c)
}
