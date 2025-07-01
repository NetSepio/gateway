package subscription

import (
	"net/http"
	"time"

	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/ctx"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/NetSepio/gateway/utils/status"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSubscription(c *gin.Context) {

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

	var input CreateSubscriptionPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		logwrapper.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := database.GetDb()

	// get plain id details by input.plainId from models.Plan using gorm
	var plan models.Plan
	if err := db.First(&plan, "id = ?", input.PlanId).Error; err != nil {
		logwrapper.Errorf("Failed to find plan with ID %s: %v", input.PlanId, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Plan not found"})
		return
	}

	// Check if there is already an active trial subscription for the user
	var existingSubscription models.SubscriptionPlan
	if err := db.Where("created_by = ? AND plan_id = ? AND status = ?", creatorId, input.PlanId, status.ACTIVE).First(&existingSubscription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// No existing subscription found, proceed to create a new one
			logwrapper.Infof("No existing subscription found for user %s with plan %s", creatorId, input.PlanId)
		} else {
			// An error occurred while checking for existing subscription
			logwrapper.Errorf("Error checking existing subscription: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "You already have an active trial subscription"})
			return
		}
	}

	if existingSubscription.CreatedAt.AddDate(0, 0, plan.Duration).After(time.Now()) && len(existingSubscription.ID) != 0 {
		logwrapper.Errorf("User %s already has an active subscription for plan %s", creatorId, plan.Name)
		// There is already an active trial subscription for the user
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already have an active trial subscription"})
		return
	}

	// Create a new trial subscription
	subscription := models.SubscriptionPlan{
		PlanID:      input.PlanId,
		Status:      status.ACTIVE,
		AutoRenewal: false,
		CreatedBy:   creatorId,
		DateCreated: time.Now(),
	}

	// Use a transaction to ensure both subscription and renewal are created atomically
	err := db.Transaction(func(tx *gorm.DB) error {
		// Save the new trial subscription to the database
		if err := tx.Model(models.SubscriptionPlan{}).Create(&subscription).Error; err != nil {
			logwrapper.Errorf("Error creating subscription: %v", err)
			return err
		}

		if plan.Duration == 0 {
			plan.Duration = 7 // default to 7 days if duration is not set
		}

		startTime := time.Now()
		endTime := startTime.AddDate(0, 0, plan.Duration)

		// save new subscription to the subscription_renewals
		subscriptionRenewal := models.SubscriptionRenewal{
			SubscriptionPlanID: subscription.ID,
			StartTime:          startTime,
			EndTime:            endTime,
			CreatedBy:          creatorId,
		}
		if err := tx.Model(models.SubscriptionRenewal{}).Create(&subscriptionRenewal).Error; err != nil {
			logwrapper.Errorf("Error creating subscription renewal: %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "subscription created"})
}
