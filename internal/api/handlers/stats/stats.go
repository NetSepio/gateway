package stats

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/stats")
	{
		g.GET("", getStats)
		g.GET("/referral", getStats) // get the user list who reffered from high to low
	}
}

func getStats(c *gin.Context) {
	db := database.GetDb()
	var queryReq GetStatsQuery
	err := c.BindQuery(&queryReq)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}
	var review []GetStatsResponse
	err = db.Model(&models.Review{}).Select("site_safety, count(site_safety)").Group("site_safety").Where(&models.Review{SiteUrl: strings.TrimSuffix(queryReq.SiteUrl, "/"), DomainAddress: queryReq.Domain}).Find(&review).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Reviews fetched successfully", review).SendD(c)
}
