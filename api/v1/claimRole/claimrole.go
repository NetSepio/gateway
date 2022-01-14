package claimrole

import (
	"net/http"

	"github.com/TheLazarusNetwork/marketplace-engine/db"
	"github.com/TheLazarusNetwork/marketplace-engine/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/cryptosign"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/claimrole")
	{
		g.Use(jwt.JWT)
		g.POST("", postClaimRole)
	}
}

func postClaimRole(c *gin.Context) {
	var req ClaimRoleRequest
	c.BindJSON(&req)

	//Message containing flowId
	role, err := getRoleByFlowId(req.FlowId)
	if err == gorm.ErrRecordNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, "flow id not found")
		return
	}
	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	message := role.Eula + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	if !isCorrect {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet address is not correct")
		return
	}

	// Update user role
	logrus.Println("walletaddress", walletAddress, "roleId", role.RoleId)
	err = db.Db.Model(&models.User{WalletAddress: walletAddress}).
		Association("Roles").
		Append(models.UserRole{WalletAddress: walletAddress, RoleId: role.RoleId}).
		Error
	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		logrus.Println(err)
		return
	} else {
		httphelper.SuccessResponse(c, "Role successfully claimed", nil)
	}

}

func getRoleByFlowId(flowId string) (models.Role, error) {
	var flowIdRecord models.FlowId
	err := db.Db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdRecord).Error
	if err != nil {
		return models.Role{}, err
	}

	var role models.Role
	err = db.Db.Model(&models.Role{}).First(&role, flowIdRecord.RelatedRoleId).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}
