package organisation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"netsepio-gateway-v1.1/internal/database"
	apikey "netsepio-gateway-v1.1/utils/apiKey"
	"netsepio-gateway-v1.1/utils/load"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/organisation")
	{
		g.POST("", createOrganisation)
		g.GET("", listOrganisations)
	}
}

// Handler: Create
func createOrganisation(c *gin.Context) {
	var input CreateOrganisationInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	load.Logger.Warn("invalid input", zap.Error(err))
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	id := uuid.New()

	if input.Name == "" {
		input.Name = "Organisation-" + id.String() // create a random name things name or noun
	}

	org := Organisation{
		ID:        id,
		Name:      input.Name,
		IPAddress: input.IPAddress,
		APIKey:    apikey.GenerateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := database.DB.Create(&org).Error; err != nil {
		load.Logger.Error("createOrganisation: DB insert failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organisation"})
		return
	}

	load.Logger.Info("Organisation created", zap.String("id", org.ID.String()))
	c.JSON(http.StatusOK, org)
}

// Handler: List all
func listOrganisations(c *gin.Context) {
	var orgs []Organisation
	if err := database.DB.Find(&orgs).Error; err != nil {
		load.Logger.Error("listOrganisations: DB fetch failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organisations"})
		return
	}
	c.JSON(http.StatusOK, orgs)
}
