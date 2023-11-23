package sotreus

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/vpn")
	{
		g.Use(paseto.PASETO)
		g.POST("", Deploy)
	}
}

func Deploy(c *gin.Context) {
	//db := dbconfig.GetDb()
	var req SotreusDeployBody
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
	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.VPN_DEPLOYER_API+"/sotreus", bytes.NewReader(ReqBodyBytes))
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

	body, _ := io.ReadAll(resp.Body)

	response := new(SotreusResponse)

	if err := json.Unmarshal(body, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// contract := models.Contract{
	// 	ContractName:    req.ContractName,
	// 	ContractAddress: response.ContractAddress,
	// 	WalletAddress:   walletAddress,
	// 	ChainId:         response.ChainId,
	// 	Verified:        response.Verified,
	// 	StorefrontId:    req.StorefrontId,
	// 	BlockNumber:     response.BlockNumber,
	// 	CollectionName:  req.CollectionName,
	// 	Thumbnail:       req.Thumbnail,
	// 	CoverImage:      req.CoverImage,
	// }
	// result := db.Create(&contract)
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	c.JSON(http.StatusOK, response)
}
