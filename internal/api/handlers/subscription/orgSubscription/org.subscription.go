package org_subscription

import (
	"net/http"
	"time"

	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/ctx"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ApplyRoutesOrgSubscriptionV11(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.Use(paseto.PASETO(false))
		g.POST("", CreateOrgSubscription)
	}
}

func CreateOrgSubscription(c *gin.Context) {

	var creatorId string

	creatorId, ctxType, err1 := ctx.GetCreatorID(c)
	if err1 != nil {
		logwrapper.Errorf("Failed to get creator ID: %v, ctxType: %s", err1, ctxType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user or organisation id not found"})
		return
	}

	if len(creatorId) == 0 {
		logwrapper.Errorf("user or organisation id not found in context")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user or organisation id not found"})
		return
	}

	var input OrgAppSubscriptionPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		logwrapper.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := database.GetDb()

	// Check if there is already an active trial subscription for the user
	var existingSubscription models.OrgSubscription

	// check active subscription from organisation subscription table
	if err := db.Where("organisation_id = ? AND status = ?", creatorId, "active").Where("? > start_time AND (? < end_time OR end_time IS NULL)", time.Now(), time.Now()).First(&existingSubscription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// No existing subscription found, proceed to create a new one
			logwrapper.Infof("No existing subscription found for organisation %s", creatorId)
		} else {
			// An error occurred while checking for existing subscription
			logwrapper.Errorf("Error checking existing subscription: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
			return
		}

		// An error occurred while checking for existing subscription
		logwrapper.Errorf("Error checking existing subscription: %v", err)
		// Save the new trial subscription to the database
		logwrapper.Infof("Creating new subscription for organisation %s ", creatorId)
		orgUUID, err := uuid.Parse(creatorId)
		if err != nil {
			logwrapper.Errorf("Invalid organisation ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organisation ID"})
			return
		}
		subscription := models.OrgSubscription{
			OrganisationID: orgUUID,
			Status:         "active",
			StartTime:      time.Now(),
			BillingCycle:   input.BillingCycle,
			AmountDue:      input.AmountDue,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		// Save the subscription to the database
		if err := db.Create(&subscription).Error; err != nil {
			logwrapper.Errorf("Failed to create subscription: %v", err)
			return
		}
		logwrapper.Infof("Subscription created for organisation %s with ID %s", creatorId, subscription.ID.String())
		c.JSON(http.StatusOK, gin.H{"status": "subscription created", "subscription_id": subscription.ID.String()})
		return
	} else {
		logwrapper.Errorf("Organisation %s already has an active subscription", creatorId)
		// There is already an active trial subscription for the organisation
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already have an active subscription"})
		return
	}
}
