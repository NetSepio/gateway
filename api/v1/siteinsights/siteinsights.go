// api/site_insight.go

package siteinsights

import (
	"errors"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/pkg/openai"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/site-insight")
	{
		g.GET("", GetSiteInsight)
	}
}

func GetSiteInsight(c *gin.Context) {
	siteURL := c.Query("siteUrl")
	if siteURL == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "siteUrl is required").SendD(c)
		return
	}

	db := dbconfig.GetDb()
	var insight models.SiteInsight
	err := db.Where("site_url = ?", siteURL).First(&insight).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			go generateAndStoreInsight(siteURL)
			httpo.NewSuccessResponse(http.StatusCreated, "Insight generation in progress").SendD(c)
			return
		} else {
			httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get insight").SendD(c)
			return
		}
	}
	httpo.NewSuccessResponseP(http.StatusOK, "Insight retrieved successfully", insight).SendD(c)
}

func generateAndStoreInsight(siteURL string) {
	content, err := ScrapeWebsiteContent(siteURL)
	if err != nil {
		logwrapper.Errorf("Error getting content: %v", err)
		return
	}
	insightText, err := openai.GenerateInsight(siteURL, content)
	if err != nil {
		logwrapper.Errorf("Error generating insight: %v", err)
		return
	}

	insight := models.SiteInsight{
		SiteURL:   siteURL,
		Insight:   insightText,
		CreatedAt: time.Now(),
	}

	db := dbconfig.GetDb()
	if err := db.Create(&insight).Error; err != nil {
		logwrapper.Errorf("Error storing insight in DB: %v", err)
	}
}
