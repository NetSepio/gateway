package organisation

import (
	"encoding/hex"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/internal/api/handlers/organisation/orgApp"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models/claims"
	apikey "github.com/NetSepio/gateway/utils/api_key"
	"github.com/NetSepio/gateway/utils/auth"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/organisation")
	{
		g.POST("", createOrganisation)
		g.GET("", listOrganisations)
		g.GET("/token", verifyOrgAPIKey)
		orgApp.ApplyRoutes(g)
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

func verifyOrgAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("X-ORG-API-KEY")
	if apiKey == "" {
		load.Logger.Warn("verifyAPIKey: API key is missing in header. Hint: Provide your organisation API key in the 'X-ORG-API-KEY' header.")
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "error": "API key is required in 'X-ORG-API-KEY' header"})
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

	// update organisation status = active
	if err := database.DB.Model(&org).Update("status", "active").Error; err != nil {
		logwrapper.Errorf("failed to update organisation status, error %v", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"valid": false, "error": "Failed to update api key status"})
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
		Token:          pasetoToken,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
