package handlers

import (
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterReferralRoutes(router *gin.Engine) {
	referral := router.Group("/referral")
	{
		referral.POST("/generate", GenerateReferralCode)
		referral.POST("/apply", ApplyReferralCode)
		referral.GET("/earnings", GetReferralEarnings)
		referral.GET("/list", GetReferrals)
	}
}

// Generate a unique referral code for a user
func GenerateReferralCode(c *gin.Context) {
	db := dbconfig.GetDb()
	userId := c.GetString("user_id") // Extract user ID from context

	// Check if the user already has a referral code
	var existingReferral models.Referral
	err := db.Where("referrer_id = ?", userId).First(&existingReferral).Error
	if err == nil {
		httpo.NewSuccessResponseP(200, "Referral code already exists", existingReferral).SendD(c)
		return
	}

	// Generate a unique referral code (e.g., first 8 chars of UUID)
	referralCode := strings.ReplaceAll(uuid.New().String(), "-", "")[:8]

	// Save referral entry
	newReferral := models.Referral{
		Id:           uuid.New().String(),
		ReferrerId:   userId,
		ReferralCode: referralCode,
	}

	if err := db.Create(&newReferral).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to create referral code").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(201, "Referral code generated successfully", newReferral).SendD(c)
}

// Apply a referral code during sign-up or payment
func ApplyReferralCode(c *gin.Context) {
	db := dbconfig.GetDb()
	var request struct {
		UserId       string `json:"userId"`
		ReferralCode string `json:"referralCode"`
	}

	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "Invalid payload").SendD(c)
		return
	}

	// Check if referral code exists
	var referrer models.Referral
	err := db.Where("referral_code = ?", request.ReferralCode).First(&referrer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "Invalid referral code").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "Database error").SendD(c)
		return
	}

	// Check if the user already used a referral code
	var existingReferral models.Referral
	err = db.Where("referee_id = ?", request.UserId).First(&existingReferral).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusConflict, "User has already used a referral code").SendD(c)
		return
	}

	// Save referral relationship
	newReferral := models.Referral{
		Id:         uuid.New().String(),
		ReferrerId: referrer.ReferrerId,
		RefereeId:  request.UserId,
	}

	if err := db.Create(&newReferral).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to apply referral code").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(201, "Referral code applied successfully", newReferral).SendD(c)
}

// Track Referral Earnings
func GetReferralEarnings(c *gin.Context) {
	db := dbconfig.GetDb()
	referrerId := c.GetString("user_id")

	var earnings []models.ReferralEarnings
	err := db.Where("referrer_id = ?", referrerId).Find(&earnings).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch earnings").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Earnings retrieved", earnings).SendD(c)
}

// List All Referrals
func GetReferrals(c *gin.Context) {
	db := dbconfig.GetDb()
	referrerId := c.GetString("user_id")

	var referrals []models.Referral
	err := db.Where("referrer_id = ?", referrerId).Find(&referrals).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to fetch referrals").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Referrals retrieved", referrals).SendD(c)
}
