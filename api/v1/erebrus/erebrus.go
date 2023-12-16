package erebrus

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
	g := r.Group("/")
	{
		g.Use(paseto.PASETO(false))
		g.POST("client/:region", RegisterClient)
	}
}
func RegisterClient(c *gin.Context) {
	region := c.Param("region")
	db := dbconfig.GetDb()
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	var count int64
	err := db.Model(&models.Erebrus{}).Where("wallet_address = ?", walletAddress).Find(&models.Erebrus{}).Count(&count).Error
	if err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
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
		UUID: "231",
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

	fmt.Println("Status:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Errorf("failed to read response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	reqBody := new(Client)
	if err := json.Unmarshal(body, reqBody); err != nil {
		logwrapper.Errorf("failed to unmarshal response: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}
	httpo.NewSuccessResponse(200, "VPN client created successfully").SendD(c)
}
