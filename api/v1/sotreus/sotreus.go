package sotreus

import (
	"bytes"
	"encoding/json"
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
	}
}

func Deploy(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var req DeployRequest
	err := c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	var DeployerRequest DeployerCreateRequest
	DeployerRequest.SotreusID = req.Name
	ReqBodyBytes, err := json.Marshal(DeployerRequest)
	if err != nil {
		logwrapper.Errorf("failed to encode request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to create VPN").SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", bytes.NewReader(ReqBodyBytes))
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

	body, _ := io.ReadAll(resp.Body)
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
	//db := dbconfig.GetDb()
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ReqBodyBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus/stop", bytes.NewReader(ReqBodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, nil)

}

func Delete(c *gin.Context) {
	db := dbconfig.GetDb()
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ReqBodyBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contractReq, err := http.NewRequest(http.MethodDelete, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", bytes.NewReader(ReqBodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		err = db.Where("name = ?", req.VpnId).Delete(&models.Sotreus{}).Error
	}
	c.JSON(http.StatusOK, nil)
}

func Start(c *gin.Context) {
	var req SotreusRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ReqBodyBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus/start", bytes.NewReader(ReqBodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(contractReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, nil)
}
