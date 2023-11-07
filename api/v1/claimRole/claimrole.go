package claimrole

import (
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/smartcontract/rawtrasaction"
	"github.com/NetSepio/gateway/generated/smartcontract/gennetsepio"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/cryptosign"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
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
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, "payload is invalid").SendD(c)
		return
	}

	//Message containing flowId
	role, err := getRoleByFlowId(req.FlowId)
	if err == gorm.ErrRecordNotFound {
		httpo.NewErrorResponse(http.StatusNotFound, "flow id not found").SendD(c)
		return
	}
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to get role by flowid, error %v", err.Error())
		return
	}
	message := fmt.Sprintf("APTOS\nmessage: %v\nnonce: %v", role.Eula, req.FlowId)
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message, req.PubKey)

	if err == cryptosign.ErrFlowIdNotFound {
		httpo.NewErrorResponse(http.StatusNotFound, err.Error()).SendD(c)
		return
	} else if err != nil {
		logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	if !isCorrect || walletAddressGin != walletAddress {
		httpo.NewErrorResponse(http.StatusForbidden, "Wallet address is not correct").SendD(c)
		return
	}

	// client := smartcontract.GetClient()
	// instance, err := netsepio.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "NETSEPIO", err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
	}
	roleIdBytesSlice, err := hexutil.Decode(role.RoleId)
	if err != nil {
		logwrapper.Warnf("failed to decode hex string : %v, for role for wallet address %v", role.RoleId, walletAddress)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}
	walletAddressHex := common.HexToAddress(walletAddress)
	var roleIdBytes [32]byte
	copy(roleIdBytes[:], roleIdBytesSlice)

	tx, err := rawtrasaction.SendRawTrasac(gennetsepio.GennetsepioABI, "grantRole", roleIdBytes, walletAddressHex)

	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Warnf("failed to grant role to user with walletaddress %v, error: %v", walletAddress, err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	err = db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{}).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to delete flowId, error %v", err.Error())
		return
	}
	payload := ClaimRolePayload{
		TransactionHash: transactionHash,
	}
	httpo.NewSuccessResponseP(200, "role grant transaction has been broadcasted", payload).SendD(c)

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
