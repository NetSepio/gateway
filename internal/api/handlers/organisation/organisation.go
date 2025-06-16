package organisation

import (
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models/claims"
	apikey "netsepio-gateway-v1.1/utils/apiKey"
	"netsepio-gateway-v1.1/utils/auth"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/organisation")
	{
		g.POST("", createOrganisation)
		g.GET("", listOrganisations)
		g.GET("/token", verifyAPIKey)
	}
}

// Handler: Create
func createOrganisation(c *gin.Context) {
	var input CreateOrganisationInput

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

func verifyAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("X-API-Key")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "error": "API key is required in 'X-API-Key' header"})
		return
	}

	var org Organisation
	if err := database.DB.Where("api_key = ?", apiKey).First(&org).Error; err != nil {
		load.Logger.Warn("verifyAPIKey: invalid API key", zap.String("api_key", apiKey))
		c.JSON(http.StatusUnauthorized, gin.H{"valid": false, "error": "Invalid API key"})
		return
	}

	customClaims := claims.NewWithOrganisation(org.ID.String(), &org.Name, &org.IPAddress)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}
	pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}

	load.Logger.Info("Organisation verified token generated sucessfully", zap.String("id", org.ID.String()))

	payload := OrganisationPaseto{
		OrganisationId: org.ID.String(),
		PasetoToken:    pasetoToken,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
