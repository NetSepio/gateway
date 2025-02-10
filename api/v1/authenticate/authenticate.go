package authenticate

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/cryptosign"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
		g.POST("/NonSign", authenticateNonSignature)
		g.Use(paseto.PASETO(false))
		g.GET("", authenticateToken)
	}
}

var CTX_CHAIN_NAME = "CHAIN_NAME"

func authenticate(c *gin.Context) {
	db := dbconfig.GetDb()
	chain_symbol := c.Query("chain") //google\
	//TODO remove flow id if 200"
	var req AuthenticateRequest

	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}

	if len(chain_symbol) == 0 && len(req.ChainName) == 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "chain name is required").SendD(c)
		return
	}

	if len(chain_symbol) != 0 {
		c.Set(CTX_CHAIN_NAME, chain_symbol)
	} else {
		c.Set(CTX_CHAIN_NAME, req.ChainName)
	}

	if len(req.ChainName) == 0 {
		req.ChainName = chain_symbol
	}

	//Get flowid type
	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", req.FlowId).First(&flowIdData).Error
	if err != nil {
		logwrapper.Errorf("failed to get flowId, error %v", err)
		httpo.NewErrorResponse(http.StatusNotFound, "flow id not found").SendD(c)
		return
	}

	if flowIdData.FlowIdType != models.AUTH {
		httpo.NewErrorResponse(http.StatusBadRequest, "flow id not created for auth").SendD(c)
		return
	}

	var isCorrect bool
	var userId string
	var walletAddr string
	if chain_symbol == "evm" {
		userAuthEULA := envconfig.EnvVars.AUTH_EULA
		message := userAuthEULA + req.FlowId
		userId, walletAddr, isCorrect, err = cryptosign.CheckSignEth(req.Signature, req.FlowId, message)

		if err == cryptosign.ErrFlowIdNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "Flow Id not found")
			return
		}

		if err != nil {
			logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	}
	if chain_symbol == "apt" {
		userAuthEULA := envconfig.EnvVars.AUTH_EULA
		message := fmt.Sprintf("APTOS\nmessage: %v\nnonce: %v", userAuthEULA, req.FlowId)

		userId, walletAddr, isCorrect, err = cryptosign.CheckSign(req.Signature, req.FlowId, message, req.PubKey)

		if err == cryptosign.ErrFlowIdNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "Flow Id not found")
			return
		}

		if err != nil {
			logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	}
	if chain_symbol == "sui" {
		userId, walletAddr, isCorrect, err = cryptosign.CheckSignSui(req.SignatureSui, req.FlowId)

		if err == cryptosign.ErrFlowIdNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "Flow Id not found")
			return
		}

		if err != nil {
			logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}

	}
	if chain_symbol == "sol" {
		walletAddr, userId, isCorrect, err = cryptosign.CheckSignSol(req.Signature, req.FlowId, req.Message, req.PubKey)

		if err == cryptosign.ErrFlowIdNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "Flow Id not found")
			return
		}

		if err != nil {
			logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			return
		}
	}
	if isCorrect {
		// update wallet address for that user_id
		err = db.Model(&models.User{}).Where("user_id = ?", userId).
			Updates(map[string]interface{}{
				"wallet_address": walletAddr,
				"chain_name":     req.ChainName,
			}).Error
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
			logwrapper.Errorf("failed to update wallet address and chain name, error %v", err.Error())
			return
		}

		customClaims := claims.NewWithWallet(userId, &walletAddr)
		pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Errorf("failed to generate token, error %v", err.Error())
			return
		}
		pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Errorf("failed to generate token, error %v", err.Error())
			return
		}
		err = db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{}).Error
		if err != nil {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Errorf("failed to delete flowId, error %v", err.Error())
			return
		}
		payload := AuthenticatePayload{
			Token:  pasetoToken,
			UserId: userId,
		}
		httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
	} else {
		httpo.NewErrorResponse(http.StatusForbidden, "Wallet Address is not correct").SendD(c)
		return
	}
}

func authenticateToken(c *gin.Context) {
	userId := c.GetString(paseto.CTX_USER_ID)
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	payload := AuthenticateTokenPayload{
		UserId:        userId,
		WalletAddress: walletAddress,
	}
	httpo.NewSuccessResponseP(200, "Token verifies successfully", payload).SendD(c)
}

func authenticateNonSignature(c *gin.Context) {
	db := dbconfig.GetDb()
	//TODO remove flow id if 200
	var req AuthenticateRequestNoSign
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}

	//Get flowid type
	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", req.FlowId).First(&flowIdData).Error
	if err != nil {
		logwrapper.Errorf("failed to get flowId, error %v", err)
		httpo.NewErrorResponse(http.StatusNotFound, "flow id not found").SendD(c)
		return
	}

	if flowIdData.FlowIdType != models.AUTH {
		httpo.NewErrorResponse(http.StatusBadRequest, "flow id not created for auth").SendD(c)
		return
	}
	if req.WalletAddress != flowIdData.WalletAddress {
		httpo.NewErrorResponse(http.StatusBadRequest, "WalletAddress incorrect").SendD(c)
		return
	}

	// update wallet address for that user_id
	err = db.Model(&models.User{}).Where("user_id = ?", flowIdData.UserId).Update("wallet_address", flowIdData.WalletAddress).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to update wallet address, error %v", err.Error())
		return
	}

	customClaims := claims.NewWithWallet(flowIdData.UserId, &flowIdData.WalletAddress)
	pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}
	pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}
	err = db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{}).Error
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to delete flowId, error %v", err.Error())
		return
	}
	payload := AuthenticatePayload{
		Token:  pasetoToken,
		UserId: flowIdData.UserId,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
