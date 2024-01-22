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
		g.GET("/clients", GetAllClients)
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
		httpo.NewErrorResponse(http.StatusBadRequest, "Can't create more clients, maximum 3 allowed").SendD(c)
		return
	}

	var req Client

	err = c.BindJSON(&req)
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}
	client := &http.Client{}
	data := Client{
		Name:         req.Name,
		Enable:       true,
		PresharedKey: req.PresharedKey,
		AllowedIPs:   []string{"0.0.0.0/0", "::/0"},
		Address:      []string{"10.0.0.0/24"},
		CreatedBy:    walletAddress,
		PublicKey:    req.PublicKey,
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

	dbEntry := models.Erebrus{
		UUID:          reqBody.Client.UUID,
		Name:          reqBody.Client.Name,
		WalletAddress: walletAddress,
		Region:        region,
		CollectionId:  req.CollectionId,
	}
	if err := db.Create(&dbEntry).Error; err != nil {
		logwrapper.Errorf("failed to create database entry: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	httpo.NewSuccessResponseP(200, "VPN client created successfully", gin.H{"client": reqBody.Client, "serverAddress": reqBody.Server.Address, "serverPublicKey": reqBody.Server.PublicKey, "endpoint": reqBody.Server.Endpoint}).SendD(c)
}

func GetClient(c *gin.Context) {
	uuid := c.Param("uuid")
	db := dbconfig.GetDb()

	var cl *models.Erebrus
	if err := db.Model(&models.Erebrus{}).Where("UUID = ?", uuid).First(&cl).Error; err != nil {
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
	err := db.Model(&models.Erebrus{}).Where("UUID = ?", uuid).First(&cl).Error
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

	if err := db.Delete(cl).Error; err != nil {
		logwrapper.Errorf("failed to delete data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "VPN client deletes successfully").SendD(c)
}

func GetConfig(c *gin.Context) {
	uuid := c.Param("uuid")
	db := dbconfig.GetDb()

	var cl *models.Erebrus
	err := db.Model(&models.Erebrus{}).Where("UUID = ?", uuid).First(&cl).Error
	if err != nil {
		logwrapper.Errorf("failed to fetch data from database: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	resp, err := http.Get(regions.ErebrusRegions[cl.Region].ServerHttp + "/api/v1.0/client/" + uuid + "/config")
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
	c.Writer.WriteHeader(200)
}

func GetClientsByRegion(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	region := c.Param("region")

	db := dbconfig.GetDb()
	var clients *[]models.Erebrus
	db.Model(&models.Erebrus{}).Where("wallet_address = ? and region = ?", walletAddress, region).Find(&clients)

	httpo.NewSuccessResponseP(200, "VPN client fetched successfully", clients).SendD(c)
}
func GetClientsByCollectionRegion(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	region := c.Param("region")
	collection_id := c.Param("collection_id")

	db := dbconfig.GetDb()
	var clients *[]models.Erebrus
	db.Model(&models.Erebrus{}).Where("wallet_address = ? and region = ? and collection_id = ?", walletAddress, region, collection_id).Find(&clients)

	httpo.NewSuccessResponseP(200, "VPN clients fetched successfully", clients).SendD(c)
}
func GetAllClients(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	region := c.Query("region")
	collectionID := c.Query("collection_id")

	db := dbconfig.GetDb()
	query := db.Model(&models.Erebrus{}).Where("wallet_address = ?", walletAddress)

	if region != "" {
		query = query.Where("region = ?", region)
	}
	if collectionID != "" {
		query = query.Where("collection_id = ?", collectionID)
	}

	var clients *[]models.Erebrus
	query.Find(&clients)

	httpo.NewSuccessResponseP(200, "VPN client fetched successfully", clients).SendD(c)
}

func GetClientsByCollectionId(c *gin.Context) {
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	collection_id := c.Param("collection_id")

	db := dbconfig.GetDb()
	var clients *[]models.Erebrus
	db.Model(&models.Erebrus{}).Where("wallet_address = ? and collection_id = ?", walletAddress, collection_id).Find(&clients)

	httpo.NewSuccessResponseP(200, "VPN clients fetched successfully", clients).SendD(c)
}
