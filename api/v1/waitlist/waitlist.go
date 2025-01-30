package waitlist

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/waitlist")
	{
		g.POST("", waitlist)
	}
}

func waitlist(c *gin.Context) {
	db := dbconfig.GetDb()
	var req WaitListRequest
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("body is invalid: %s", err)).SendD(c)
		return
	}

	findResult := db.Model(&models.WaitList{}).Find(&models.WaitList{}, &models.WaitList{EmailId: req.EmailId})

	if err := findResult.Error; err != nil {
		logwrapper.Errorf("failed to check if user exist in waitlist, %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	if findResult.RowsAffected > 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "Already exist in waitlist").SendD(c)
		return
	}

	newWailListMember := &models.WaitList{
		EmailId:       req.EmailId,
		WalletAddress: strings.ToLower(req.WalletAddress),
		Twitter:       req.Twitter,
		Discord:       req.Discord,
	}
	if err := db.Create(newWailListMember).Error; err != nil {
		logwrapper.Errorf("failed to add to waitlist, %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	if err != nil {
		logwrapper.Errorf("failed to add to waitlist, %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}
	httpo.NewSuccessResponse(200, "Added in waitlist").SendD(c)
}
