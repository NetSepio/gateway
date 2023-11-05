package waitlist

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
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
		httphelper.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("body is invalid: %s", err))
		return
	}

	findResult := db.Model(&models.WaitList{}).Find(&models.WaitList{}, &models.WaitList{EmailId: req.EmailId})

	if err := findResult.Error; err != nil {
		logwrapper.Errorf("failed to check if user exist in waitlist, %s", err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	if findResult.RowsAffected > 0 {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Already exist in waitlist")
		return
	}

	newWailListMember := &models.WaitList{
		EmailId:       req.EmailId,
		WalletAddress: req.WalletAddress,
		Twitter:       req.Twitter,
	}
	if err := db.Create(newWailListMember).Error; err != nil {
		logwrapper.Errorf("failed to add to waitlist, %s", err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	if err != nil {
		logwrapper.Errorf("failed to add to waitlist, %s", err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Added in waitlist", nil)
}
