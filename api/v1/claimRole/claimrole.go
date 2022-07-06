package claimrole

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/smartcontract/rawtrasaction"
	"github.com/NetSepio/gateway/generated/smartcontract/gennetsepio"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/cryptosign"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/claimrole")
	{
		g.Use(paseto.PASETO)
		g.POST("", postClaimRole)
	}
}

func postClaimRole(c *gin.Context) {
	walletAddressGin := c.GetString("walletAddress")
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

	if !isCorrect || walletAddressGin != walletAddress {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet address is not correct")
		return
	}

	// client := smartcontract.GetClient()
	// instance, err := netsepio.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "NETSEPIO", err.Error())
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

	tx, err := rawtrasaction.SendRawTrasac(gennetsepio.GennetsepioABI, "grantRole", roleIdBytes, walletAddressHex)

	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		logwrapper.Warnf("failed to grant role to user with walletaddress %v, error: %v", walletAddress, err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{})
	payload := ClaimRolePayload{
		TransactionHash: transactionHash,
	}
	httphelper.SuccessResponse(c, "Role successfully claimed", payload)

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
