package referral

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
)

func ApplyReferraSubscriptionlRoutes(r *gin.RouterGroup) {
	r.Use(paseto.PASETO(true))
	referral := r.Group("/referral/subscription")
	{
		referral.GET("/earnings", GetReferralEarnings)
		referral.PATCH("", ApplyReferralCodeForAccount)
	}
}

// Track Referral Earnings
func GetReferralEarnings(c *gin.Context) {
	db := database.GetDb()
	referrerId := c.GetString("user_id")

	var earnings []models.ReferralEarnings
	err := db.Where("referrer_id = ?", referrerId).Find(&earnings).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch earnings").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Earnings retrieved", earnings).SendD(c)
}

func ApplyReferralCodeForSubscription(c *gin.Context) {

	var request struct {
		ReferralCode string `json:"referralCode"`
	}

	// get userid from the paseto token
	userId := c.GetString(paseto.CTX_USER_ID)

	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "Invalid payload").SendD(c)
		return
	}

	db := database.GetDb()

	tx := db.Begin()

	// Check if the user already used a referral code
	var existingReferral models.ReferralSubscription

	err := db.Where("referee_id = ? AND referral_code = ?", userId, request.ReferralCode).First(&existingReferral).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusConflict, "User has already used this referral code").SendD(c)
		return
	}

	err = db.Where("referee_id = ? AND referral_code = ?", userId, request.ReferralCode).First(&existingReferral).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusConflict, "User has already used this referral code").SendD(c)
		return
	}

	// Update the existing referral record with the new referee
	err = tx.Model(&models.ReferralAccount{}).
		Where("referral_code = ?", request.ReferralCode).
		Update("referred_id", userId).Error

	if err != nil {
		logrus.Errorln("Failed to Update the existing referral record with the new referee: ", err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to update referral data -> "+err.Error()).SendD(c)
		tx.Rollback()
		return
	}

	// Check if referral code exists in user table by referral code
	var user models.User
	err = db.Where("referral_code = ?", request.ReferralCode).First(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Errorln("No Record Found")
			httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve user data").SendD(c)
			return

		}
		logrus.Errorln("Functional Error -> ApplyReferralCode: " + err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	referralEarnings := models.ReferralEarnings{
		Id:           uuid.New().String(),
		ReferrerId:   user.UserId,
		ReferredId:   userId,
		AmountEarned: 10.0,
	}
	if err := tx.Create(&referralEarnings).Error; err != nil {
		logrus.Errorln("Failed to create referral earnings record: ", err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create referral earnings").SendD(c)
		tx.Rollback()
		return
	}

	tx.Commit()

	httpo.NewSuccessResponseP(http.StatusOK, "Referral code applied successfully", struct{}{}).SendD(c)
}
