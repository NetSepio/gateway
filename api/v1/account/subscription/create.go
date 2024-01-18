package subscription

import (
	"math"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/aptos"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
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
	// update the user's stripePiId
	err = db.Model(&models.User{}).Where("user_id = ?", userId).Update("stripe_pi_id", pi.ID).Error
	if err != nil {
		logwrapper.Errorf("failed to update stripe_pi_id: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(http.StatusOK, "payment intent created", BuyErebrusNFTResponse{ClientSecret: pi.ClientSecret}).SendD(c)
}
