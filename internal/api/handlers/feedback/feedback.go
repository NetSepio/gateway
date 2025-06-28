package feedback

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/feedback")
	{
		g.Use(paseto.PASETO(false))
		g.POST("", createFeedback)
		g.GET("", getFeedback)
	}
}

func createFeedback(c *gin.Context) {
	db := database.GetDb()
	var newFeedback models.UserFeedback
	err := c.BindJSON(&newFeedback)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	newFeedback.UserId = c.GetString(paseto.CTX_USER_ID)

	// Check if the user's feedback already exists
	var existingFeedback models.UserFeedback
	result := db.Model(&models.UserFeedback{}).Where("user_id = ?", newFeedback.UserId).First(&existingFeedback)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logwrapper.Errorf("failed to check existing feedback: %s", result.Error)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to check existing feedback").SendD(c)
		return
	}

	// If the user's feedback exists, update it; otherwise, create a new entry
	if result.RowsAffected > 0 {
		existingFeedback.CreatedAt = time.Now() // Optionally update the timestamp
		existingFeedback.Feedback = newFeedback.Feedback
		existingFeedback.Rating = newFeedback.Rating
		if err := db.Model(&models.UserFeedback{}).Where("user_id = ?", newFeedback.UserId).Updates(&existingFeedback).Error; err != nil {
			logwrapper.Errorf("failed to update existing feedback: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to update existing feedback").SendD(c)
			return
		}
		httpo.NewSuccessResponse(http.StatusOK, "Feedback updated successfully").SendD(c)
	} else {
		if err := db.Create(&newFeedback).Error; err != nil {
			logwrapper.Errorf("failed to add new feedback: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to add new feedback").SendD(c)
			return
		}
		httpo.NewSuccessResponse(http.StatusOK, "Feedback added successfully").SendD(c)
	}
}

func getFeedback(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	db := database.GetDb()
	var userFeedback models.UserFeedback
	if err := db.Model(&models.UserFeedback{}).Where("user_id = ?", userId).First(&userFeedback).Error; err != nil {
		logwrapper.Errorf("failed to retrieve user feedback: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve user feedback").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(http.StatusOK, "User feedback retrieved successfully", userFeedback).SendD(c)
}
