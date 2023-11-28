package sotreus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/vpn")
	{
		g.Use(paseto.PASETO)
		g.POST("", Deploy)
		g.POST("/stop", Stop)
		g.DELETE("", Delete)
		g.POST("/start", Start)
		g.GET("/all", AllDeployments)
		g.GET("/", MyDeployments)
	}
}

func Deploy(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	var req DeployRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	deployerRequest := DeployerCreateRequest{SotreusID: req.Name, WalletAddress: walletAddress}
	reqBodyBytes, err := json.Marshal(deployerRequest)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", bytes.NewReader(reqBodyBytes))
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	response := new(SotreusResponse)

	if err := json.Unmarshal(body, response); err != nil {
		logwrapper.Errorf("failed to get response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contract := models.Sotreus{
		Name:          response.Message.VpnID,
		WalletAddress: walletAddress,
		Region:        req.Region,
	}
	result := db.Create(&contract)
	if result.Error != nil {
		logwrapper.Errorf("failed to create db entry: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	payload := DeployResponse{
		VpnID:             response.Message.VpnID,
		VpnEndpoint:       response.Message.VpnEndpoint,
		FirewallEndpoint:  response.Message.FirewallEndpoint,
		DashboardPassword: response.Message.DashboardPassword,
	}
	httpo.NewSuccessResponseP(200, "VPN deployment successful", payload).SendD(c)
}

func Stop(c *gin.Context) {
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus/stop", bytes.NewReader(reqBodyBytes))
	if err != nil {
		logwrapper.Errorf("failed to create request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()

	httpo.NewSuccessResponse(200, "VPN deployment stopped").SendD(c)
}

func Delete(c *gin.Context) {
	db := dbconfig.GetDb()
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}

	ReqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodDelete, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", bytes.NewReader(ReqBodyBytes))
	if err != nil {
		logwrapper.Errorf("failed to create request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		err = db.Where("name = ?", req.VpnId).Delete(&models.Sotreus{}).Error
		if err != nil {
			logwrapper.Errorf("failed to create DB entry: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
			return
		}
	}
	httpo.NewSuccessResponse(200, "VPN deployment deleted").SendD(c)
}

func Start(c *gin.Context) {
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}

	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus/start", bytes.NewReader(reqBodyBytes))
	if err != nil {
		logwrapper.Errorf("failed to create request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()

	httpo.NewSuccessResponse(200, "VPN deployment started").SendD(c)
}

func MyDeployments(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	contractReq, err := http.NewRequest(http.MethodGet, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus/"+walletAddress, nil)
	if err != nil {
		logwrapper.Errorf("failed to create request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to read response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	response := new(GetDeployments)

	if err := json.Unmarshal(body, response); err != nil {
		logwrapper.Errorf("failed to decode response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Fetched Deployments", response.Data).SendD(c)
}
func AllDeployments(c *gin.Context) {
	contractReq, err := http.NewRequest(http.MethodGet, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", nil)
	if err != nil {
		logwrapper.Errorf("failed to create request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to send request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to read response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	response := new(GetDeployments)

	if err := json.Unmarshal(body, response); err != nil {
		logwrapper.Errorf("failed to decode response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Fetched all deployments", response.Data).SendD(c)
}
