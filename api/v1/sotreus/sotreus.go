package sotreus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/constants/regions"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/vpn")
	{
		g.Use(paseto.PASETO(false))
		g.POST("", Deploy)
		g.POST("/stop", Stop)
		g.DELETE("", Delete)
		g.POST("/start", Start)
		g.GET("/all/:region", AllDeployments)
		g.GET("/:region", MyDeployments)
	}
}

func Deploy(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	var count int64
	err := db.Model(&models.Sotreus{}).Where("wallet_address = ?", walletAddress).Find(&models.Sotreus{}).Count(&count).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	if count >= 1 {
		logwrapper.Error("Can't create more vpn instances, maximum 1 allowed")
		httpo.NewErrorResponse(http.StatusBadRequest, "Can't create more vpn instances, maximum 1 allowed").SendD(c)
		return
	}

	var req DeployRequest
	err = c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	ServerLink := regions.Regions[req.Region].ServerHttp
	deployerRequest := DeployerCreateRequest{SotreusID: req.Name, WalletAddress: walletAddress, Region: regions.Regions[req.Region].Code}
	reqBodyBytes, err := json.Marshal(deployerRequest)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, ServerLink+"/sotreus", bytes.NewReader(reqBodyBytes))
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
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
	db := dbconfig.GetDb()
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	var vpn models.Sotreus
	err = db.Model(&models.Sotreus{}).Where("name = ?", req.VpnId).First(&vpn).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch data from database").SendD(c)
		return
	}
	ServerLink := regions.Regions[vpn.Region].ServerHttp
	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, ServerLink+"/sotreus/stop", bytes.NewReader(reqBodyBytes))
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
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
	var vpn models.Sotreus
	err = db.Model(&models.Sotreus{}).Where("name = ?", req.VpnId).First(&vpn).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch data from database").SendD(c)
		return
	}
	ServerLink := regions.Regions[vpn.Region].ServerHttp
	ReqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodDelete, ServerLink+"/sotreus", bytes.NewReader(ReqBodyBytes))
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
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
	err = db.Model(&models.Sotreus{}).Where("name = ?", vpn.Name).Delete(&models.Sotreus{}).Error
	if err != nil {
		logwrapper.Errorf("failed to delete vpn from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to delete vpn from database").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "VPN deployment deleted").SendD(c)
}

func Start(c *gin.Context) {
	db := dbconfig.GetDb()
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}
	var vpn models.Sotreus
	err = db.Model(&models.Sotreus{}).Where("name = ?", req.VpnId).First(&vpn).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch data from database").SendD(c)
		return
	}
	ServerLink := regions.Regions[vpn.Region].ServerHttp
	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, ServerLink+"/sotreus/start", bytes.NewReader(reqBodyBytes))
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
		return
	}
	defer resp.Body.Close()

	httpo.NewSuccessResponse(200, "VPN deployment started").SendD(c)
}

func MyDeployments(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	region := c.Param("region")
	ServerLink := regions.Regions[region].ServerHttp
	contractReq, err := http.NewRequest(http.MethodGet, ServerLink+"/sotreus/"+walletAddress, nil)
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
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
	region := c.Param("region")
	ServerLink := regions.Regions[region].ServerHttp
	contractReq, err := http.NewRequest(http.MethodGet, ServerLink+"/sotreus", nil)
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
	if resp.StatusCode != 200 {
		logwrapper.Errorf("Error in response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Error in response").SendD(c)
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
