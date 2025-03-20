package referral

import (
	"net/http"

	"github.com/NetSepio/gateway/models"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up API routes
func ApplyRoutes(router *gin.Engine) {
	referralRoutes := router.Group("/referral-discounts")
	{
		referralRoutes.POST("/", CreateReferralDiscountHandler)
		referralRoutes.GET("/", GetAllReferralDiscountsHandler)
		referralRoutes.GET("/:id", GetReferralDiscountHandler)
		referralRoutes.PUT("/:id", UpdateReferralDiscountHandler)
		referralRoutes.DELETE("/:id", DeleteReferralDiscountHandler)
	}
}

// CreateReferralDiscountHandler handles creating a referral discount
func CreateReferralDiscountHandler(c *gin.Context) {
	var referral models.ReferralDiscount

	if err := c.ShouldBindJSON(&referral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateReferralDiscount(&referral); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create referral discount"})
		return
	}

	c.JSON(http.StatusCreated, referral)
}

// GetReferralDiscountHandler retrieves a referral discount by ID
func GetReferralDiscountHandler(c *gin.Context) {
	id := c.Param("id")
	referral, err := GetReferralDiscount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Referral discount not found"})
		return
	}
	c.JSON(http.StatusOK, referral)
}

// GetAllReferralDiscountsHandler retrieves all referral discounts
func GetAllReferralDiscountsHandler(c *gin.Context) {
	referrals, err := GetAllReferralDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve referral discounts"})
		return
	}
	c.JSON(http.StatusOK, referrals)
}

// UpdateReferralDiscountHandler updates an existing referral discount
func UpdateReferralDiscountHandler(c *gin.Context) {
	var referral models.ReferralDiscount

	if err := c.ShouldBindJSON(&referral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateReferralDiscount(&referral); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update referral discount"})
		return
	}

	c.JSON(http.StatusOK, referral)
}

// DeleteReferralDiscountHandler deletes a referral discount by ID
func DeleteReferralDiscountHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteReferralDiscount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete referral discount"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Referral discount deleted successfully"})
}
