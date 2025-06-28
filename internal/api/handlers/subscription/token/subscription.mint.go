package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
)

func ApplyRoutesSubscriptionMint(r *gin.RouterGroup) {
	g := r.Group("/mint")
	{
		g.POST("", createMintAddress)
		g.GET("/all", getMintAddresses)
		g.GET("/:id", getMintAddressByID)
		g.DELETE("/:id", deleteMintAddress)

	}
}

func createMintAddress(c *gin.Context) {
	var mint models.NFTSubscriptionMintAddress
	if err := c.ShouldBindJSON(&mint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB2()

	mint.ID = uuid.New()
	if err := db.Create(&mint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mint)
}

func getMintAddresses(c *gin.Context) {
	var mints []models.NFTSubscriptionMintAddress
	db := database.GetDB2()
	tx := db.Find(&mints)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, mints)
}

func getMintAddressByID(c *gin.Context) {
	id := c.Param("id")
	var mint models.NFTSubscriptionMintAddress
	db := database.GetDB2()

	if err := db.First(&mint, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mint)
}

func deleteMintAddress(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB2()
	if err := db.Delete(&models.NFTSubscriptionMintAddress{}, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
