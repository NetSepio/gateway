package claimrole

import (
	"net/http"

	"github.com/TheLazarusNetwork/marketplace-engine/api/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract/rawtrasaction"
	gcreatify "github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/cryptosign"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	db := dbconfig.GetDb()
	var req ClaimRoleRequest
	c.BindJSON(&req)

	//Message containing flowId
	role, err := getRoleByFlowId(req.FlowId)
	if err == gorm.ErrRecordNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, "flow id not found")
		return
	}
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get role by flowid, error %v", err.Error())
		return
	}
	message := role.Eula + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	if !isCorrect {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet address is not correct")
		return
	}

	// client := smartcontract.GetClient()
	// instance, err := creatify.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "CREATIFY", err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
	}
	roleIdBytesSlice, err := hexutil.Decode(role.RoleId)
	if err != nil {
		logwrapper.Warnf("failed to decode hex string : %v, for role for wallet address %v", role.RoleId, walletAddress)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	walletAddressHex := common.HexToAddress(walletAddress)
	var roleIdBytes [32]byte
	copy(roleIdBytes[:], roleIdBytesSlice)
	if err != nil {
		logwrapper.Errorf("failed to parse ABI for %v, error: %v", "CREATIFY", err.Error())
		httphelper.ErrResponse(c, 500, "unexpected error occured")
		return
	}
	// authBindOpts, err := auth.GetAuth(client)

	if err != nil {
		logwrapper.Errorf("failed to get auth, error: %v", err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	tx, err := rawtrasaction.SendRawTrasac(gcreatify.CreatifyABI, "grantRole", roleIdBytes, walletAddressHex)

	// tx, err := instance.GrantRole(authBindOpts, roleIdBytes, walletAddressHex)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		logwrapper.Warnf("failed to grant role to user with walletaddress %v, error: %v", walletAddress, err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	// Update user role
	err = db.Model(&models.User{WalletAddress: walletAddress}).
		Association("Roles").
		Append(models.UserRole{WalletAddress: walletAddress, RoleId: role.RoleId}).
		Error
	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		logwrapper.Error(err)
		return
	} else {
		db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{})
		payload := ClaimRolePayload{
			TransactionHash: transactionHash,
		}
		httphelper.SuccessResponse(c, "Role successfully claimed", payload)
	}

}

func getRoleByFlowId(flowId string) (models.Role, error) {
	db := dbconfig.GetDb()
	var flowIdRecord models.FlowId
	err := db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdRecord).Error
	if err != nil {
		return models.Role{}, err
	}

	var role models.Role
	err = db.Model(&models.Role{}).Where("role_id = ?", flowIdRecord.RelatedRoleId).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}
