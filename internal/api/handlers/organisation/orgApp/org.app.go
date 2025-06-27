package orgApp

import (
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/utils/api_key"
	"github.com/NetSepio/gateway/utils/auth"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
)

// Input struct for create/update
type OrganisationAppInput struct {
	OrganisationId uuid.UUID `json:"organisation_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
}

// Register routes
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/app")
	{
		g.POST("", createOrganisationApp)
		g.GET("", listOrganisationApps)
		g.GET("/:id", getOrganisationApp)
		g.PUT("/:id", updateOrganisationApp)
		g.DELETE("/:id", deleteOrganisationApp)
		g.GET("/token", verifyAppAPIKey)
	}
}

// Create
func createOrganisationApp(c *gin.Context) {
	var input OrganisationAppInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()
	if input.Name == "" {
		input.Name = "App-" + id.String()
	}
	if input.OrganisationId == uuid.Nil || len(input.OrganisationId) == 0 {
		// logging the error
		load.Logger.Sugar().Info("createOrganisationApp: Creating an organisation for app")
		// create organtisation
		orgId := uuid.New()

		org := models.Organisation{
			ID:        orgId,
			Name:      "App-" + orgId.String(),
			APIKey:    "App based",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := database.DB.Create(&org).Error; err != nil {
			load.Logger.Error("createOrganisation: DB insert failed", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organisation"})
			return
		}
		input.OrganisationId = org.ID
		load.Logger.Info("Organisation created for app", zap.String("id", org.ID.String()))
	} else {
		// verify this organisation exists
		var org models.Organisation
		if err := database.DB.First(&org, "id = ?", input.OrganisationId).Error; err != nil {
			load.Logger.Error("createOrganisationApp: Organisation not found", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid organisation"})
			return
		}
	}

	app := OrganisationApp{
		ID:             id,
		OrganisationId: input.OrganisationId,
		Name:           input.Name,
		Description:    input.Description,
		APIKey:         api_key.GenerateAPIKey(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := database.DB.Create(&app).Error; err != nil {
		load.Logger.Error("createOrganisationApp: DB insert failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organisation app"})
		return
	}

	load.Logger.Info("OrganisationApp created", zap.String("id", app.ID.String()))
	c.JSON(http.StatusOK, app)
}

// List all
func listOrganisationApps(c *gin.Context) {
	var apps []OrganisationApp
	if err := database.DB.Find(&apps).Error; err != nil {
		load.Logger.Error("listOrganisationApps: DB fetch failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organisation apps"})
		return
	}
	c.JSON(http.StatusOK, apps)
}

// Get by ID
func getOrganisationApp(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var app OrganisationApp
	if err := database.DB.First(&app, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organisation app not found"})
		return
	}
	c.JSON(http.StatusOK, app)
}

// Update
func updateOrganisationApp(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input OrganisationAppInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var app OrganisationApp
	if err := database.DB.First(&app, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organisation app not found"})
		return
	}

	app.Name = input.Name
	app.Description = input.Description
	app.OrganisationId = input.OrganisationId
	app.UpdatedAt = time.Now()

	if err := database.DB.Save(&app).Error; err != nil {
		load.Logger.Error("updateOrganisationApp: DB update failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organisation app"})
		return
	}

	c.JSON(http.StatusOK, app)
}

// Delete
func deleteOrganisationApp(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := database.DB.Delete(&OrganisationApp{}, "id = ?", id).Error; err != nil {
		load.Logger.Error("deleteOrganisationApp: DB delete failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organisation app"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func verifyAppAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("X-APP-API-KEY")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"valid": false, "error": "API key is required in 'X-API-Key' header"})
		return
	}

	var orgApp models.OrganisationApp
	if err := database.DB.Where("api_key = ?", apiKey).First(&orgApp).Error; err != nil {
		load.Logger.Warn("verifyAPIKey: invalid API key", zap.String("api_key", apiKey))
		c.JSON(http.StatusUnauthorized, gin.H{"valid": false, "error": "Invalid API key"})
		return
	}

	customClaims := claims.NewWithOrganisationApp(orgApp.OrganisationId.String(), orgApp.ID.String(), &orgApp.Name)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}

	// update orgAppanisation status = active
	if err := database.DB.Model(&orgApp).Update("status", "active").Error; err != nil {
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

	load.Logger.Info("Organisation verified token generated sucessfully", zap.String("id", orgApp.ID.String()))

	payload := OrganisationAppPaseto{
		OrganisationId: orgApp.ID.String(),
		Token:          pasetoToken,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
