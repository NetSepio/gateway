package stats

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/stats")
	{
		g.Use(paseto.PASETO)
		g.GET("", getStats)
	}
}

func getStats(c *gin.Context) {
	db := dbconfig.GetDb()
	var queryReq GetStatsQuery
	err := c.BindQuery(&queryReq)
	if err != nil {
		//TODO not override status or not set status again
		httphelper.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err))
		return
	}
	var review []GetStatsResponse
	err = db.Model(&models.Review{}).Select("site_safety, count(site_safety)").Group("site_safety").Where(&models.Review{SiteUrl: queryReq.SiteUrl, DomainAddress: queryReq.Domain}).Find(&review).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Reviews fetched successfully", review)
}
