// api/site_insight.go

package siteinsights

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/pkg/openai"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetSiteInsightQuery struct {
	SiteURL string `form:"siteUrl" binding:"required,http_url"`
}

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/site-insight")
	{
		g.GET("", GetSiteInsight)
	}
}

func GetSiteInsight(c *gin.Context) {
	var query GetSiteInsightQuery
	err := c.BindQuery(&query)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("failed to validate request: %s", err)).SendD(c)
		return
	}

	db := dbconfig.GetDb()
	var insight models.SiteInsight
	err = db.Where("site_url = ?", query.SiteURL).First(&insight).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			go generateAndStoreInsight(query.SiteURL)
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
