package erebrus

import (
	"bytes"
	"encoding/json"
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
	g := r.Group("/erebrus")
	{
		g.Use(paseto.PASETO(false))
		g.POST("/client/:region", RegisterClient)
		g.GET("/client/:region/:uuid", GetClient)
		g.GET("/client/:region", GetClients)
		g.DELETE("/client/:region/:uuid", DeleteClient)
		g.GET("/config/:region/:uuid", GetConfig)
	}
}
func RegisterClient(c *gin.Context) {
	region := c.Param("region")
	db := dbconfig.GetDb()
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	var count int64
	err := db.Model(&models.Erebrus{}).Where("wallet_address = ?", walletAddress).Find(&models.Erebrus{}).Count(&count).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	if count >= 3 {
		logwrapper.Error("Can't create more clients, maximum 3 allowed")
		httpo.NewErrorResponse(http.StatusInternalServerError, "Can't create more clients, maximum 3 allowed").SendD(c)
		return
	}

	var req Client

	err = c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	client := &http.Client{}
	data := Client{
		Name:       req.Name,
		Enable:     true,
		AllowedIPs: []string{"0.0.0.0/0", "::/0"},
		Address:    []string{"10.0.0.0/24"},
		CreatedBy:  walletAddress,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logwrapper.Errorf("failed to Marshal data: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, regions.ErebrusRegions[region].ServerHttp+"/api/v1.0/client", bytes.NewReader(dataBytes))
	if err != nil {
		logwrapper.Errorf("failed to create	 request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to perform request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to read response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	reqBody := new(Response)
	if err := json.Unmarshal(body, reqBody); err != nil {
		logwrapper.Errorf("failed to unmarshal response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	httpo.NewSuccessResponseP(200, "VPN client created successfully", reqBody.Client).SendD(c)
}

func GetClient(c *gin.Context) {
	uuid := c.Param("uuid")
	db := dbconfig.GetDb()

	var cl *models.Erebrus
	if err := db.First(&cl, uuid).Error; err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	resp, err := http.Get(regions.ErebrusRegions[cl.Region].ServerHttp + "/api/v1.0/client/" + uuid)
	if err != nil {
		logwrapper.Errorf("failed to create	 request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to read response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	resBody := new(Response)
	if err := json.Unmarshal(body, resBody); err != nil {
		logwrapper.Errorf("failed to unmarshal response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	httpo.NewSuccessResponseP(200, "VPN client fetched successfully", resBody.Client).SendD(c)
}

func DeleteClient(c *gin.Context) {
	uuid := c.Param("uuid")
	db := dbconfig.GetDb()

	var cl *models.Erebrus
	err := db.First(&cl, uuid).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	client := &http.Client{}
	contractReq, err := http.NewRequest(http.MethodDelete, regions.ErebrusRegions[cl.Region].ServerHttp+"/api/v1.0/client", bytes.NewReader(nil))
	if err != nil {
		logwrapper.Errorf("failed to create	 request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	resp, err := client.Do(contractReq)
	if err != nil {
		logwrapper.Errorf("failed to perform request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	defer resp.Body.Close()

	httpo.NewSuccessResponse(200, "VPN client deletes successfully").SendD(c)
}

func GetConfig(c *gin.Context) {
	uuid := c.Param("uuid")
	db := dbconfig.GetDb()

	var cl *models.Erebrus
	err := db.First(&cl, uuid).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	resp, err := http.Get(regions.ErebrusRegions[cl.Region].ServerHttp + "/api/v1.0/server/config")
	if err != nil {
		logwrapper.Errorf("failed to create	request: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	defer resp.Body.Close()

	c.Header("Content-Disposition", "attachment; filename="+cl.Name+".conf")
	c.Header("Content-Type", resp.Header.Get("Content-Type"))

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	httpo.NewSuccessResponse(200, "VPN config fetched successfully").SendD(c)
}

func GetClients(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	db := dbconfig.GetDb()
	var clients *[]models.Erebrus
	db.Model(&models.Erebrus{}).Where("wallet_address = ?", walletAddress).Find(&clients)

	httpo.NewSuccessResponseP(200, "VPN client fetched successfully", clients).SendD(c)
}
