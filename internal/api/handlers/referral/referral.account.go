package referral

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ApplyReferraAccountlRoutes(r *gin.RouterGroup) {
	r.Use(paseto.PASETO(true))
	referral := r.Group("/referral/account")
	{
		referral.GET("", GetReferrals)
		referral.POST("", GenerateReferralCode) // TODO if the referal code exist do not generate
		referral.PATCH("", ApplyReferralCodeForAccount)
	}
}

// Generate a unique referral code for a user
func GenerateReferralCode(c *gin.Context) {
	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID) // Extract user ID from context

	// Check if the user already has a referral code
	var existingReferral models.ReferralAccount
	err := db.Where("referrer_id = ? and referee_id = ?", userId, userId).First(&existingReferral).Error
	if err == nil {
		httpo.NewSuccessResponseP(200, "Referral code already exists", existingReferral).SendD(c)
		return
	}

	// Generate a unique referral code (e.g., first 8 chars of UUID)
	referralCode := strings.ReplaceAll(uuid.New().String(), "-", "")[:8]

	// Save referral entry
	newReferral := models.ReferralAccount{
		Id:           uuid.New().String(),
		ReferrerId:   userId,
		ReferredId:   uuid.Nil.String(),
		ReferralCode: referralCode,
	}

	fmt.Println()
	fmt.Printf("%+v\n", newReferral)
	fmt.Println()

	if err := db.Create(&newReferral).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create referral code").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(201, "Referral code generated successfully", newReferral).SendD(c)
}
func GenerateReferralCodeForAnotherUser(c *gin.Context) {
	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID) // Extract user ID from context

	// Check if the user already has a referral code
	var existingReferral models.ReferralAccount
	err := db.Where("referrer_id = ? and referee_id = ?", userId, userId).First(&existingReferral).Error
	if err == nil {
		httpo.NewSuccessResponseP(200, "Referral code already exists", existingReferral).SendD(c)
		return
	}

	// Generate a unique referral code (e.g., first 8 chars of UUID)
	referralCode := strings.ReplaceAll(uuid.New().String(), "-", "")[:8]

	// Save referral entry
	newReferral := models.ReferralAccount{
		Id:           uuid.New().String(),
		ReferrerId:   userId,
		ReferredId:   userId,
		ReferralCode: referralCode,
	}

	fmt.Println()
	fmt.Printf("%+v\n", newReferral)
	fmt.Println()

	if err := db.Create(&newReferral).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create referral code").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(201, "Referral code generated successfully", newReferral).SendD(c)
}

// Apply a referral code during sign-up or payment
func ApplyReferralCodeForAccount(c *gin.Context) {

	var request struct {
		ReferralCode string `json:"referralCode"`
	}

	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "Invalid payload").SendD(c)
		return
	}

	db := database.GetDb()
	// Check if referral code exists in user table by referral code
	var user models.User
	err := db.Where("referral_code = ?", request.ReferralCode).First(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Errorln("Functional Error -> ApplyReferralCodeForAccount: No Record Found For This Refferral Code")
			httpo.NewErrorResponse(http.StatusInternalServerError, "Invalid Referral Code").SendD(c)
			return

		}
		logrus.Errorln("Functional Error -> ApplyReferralCode: " + err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	// get userid from the paseto token
	userId := c.GetString(paseto.CTX_USER_ID)

	// Check if the user already used a referral code
	var existingReferral models.ReferralAccount

	err = db.Where("refereed_id = ? AND referral_code = ?", userId, request.ReferralCode).First(&existingReferral).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusConflict, "User has already used this referral code").SendD(c)
		return
	}

	// check if the user can use the referral code to whome he has referred
	err = db.Where("refereer_id = ? AND refereed_id = ?", user.UserId, userId).First(&existingReferral).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusConflict, "You cannot use the referral code of a user you have already referred.").SendD(c)
		return
	}

	newReferral := models.ReferralAccount{
		Id:           uuid.NewString(),
		ReferrerId:   user.UserId, // Assign the referrer directly from the request
		ReferredId:   userId,
		ReferralCode: request.ReferralCode,
		CreatedAt:    time.Now(),
	}

	err = db.Create(&newReferral).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"referrer_id":   user.UserId,
			"referred_id":   userId,
			"referral_code": request.ReferralCode,
			"error":         err.Error(),
		}).Error("Failed to insert referral record")
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to update referral data -> "+err.Error()).SendD(c)
		return
	}

	httpo.NewSuccessResponseP(http.StatusOK, "Referral code applied successfully", struct{}{}).SendD(c)
}

// List All Referrals
func GetReferrals(c *gin.Context) {
	db := database.GetDb()
	referrerId := c.GetString("user_id")

	var referrals []models.ReferralAccount
	err := db.Where("referrer_id = ?", referrerId).Find(&referrals).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch referrals").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Referrals retrieved", referrals).SendD(c)
}

func GenerateReferralCodeForUser(user models.User) string {
	db := database.GetDb()
	// Generate a unique referral code (e.g., first 8 chars of UUID)
	referralCode := strings.ReplaceAll(uuid.New().String(), "-", "")[:8]
	// Save referral entry in user table
	user.ReferralCode = referralCode
	err := db.Save(&user).Error
	if err != nil {
		logrus.Errorf("failed to save referral code, error %v", err)
		return ""
	}
	return referralCode
}

func GetReferalCode() (referralCode string) {
	// Generate a unique referral code (e.g., first 8 chars of UUID)
	referralCode = strings.ReplaceAll(uuid.New().String(), "-", "")[:8]
	return
}
