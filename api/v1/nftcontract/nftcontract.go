package nftcontract

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const NFTContractABI = `[
    {"inputs": [], "name": "name", "outputs": [{"type": "string"}], "stateMutability": "view", "type": "function"},
    {"inputs": [], "name": "symbol", "outputs": [{"type": "string"}], "stateMutability": "view", "type": "function"},
    {"inputs": [], "name": "totalSupply", "outputs": [{"type": "uint256"}], "stateMutability": "view", "type": "function"},
    {"inputs": [], "name": "owner", "outputs": [{"type": "address"}], "stateMutability": "view", "type": "function"},
    {"inputs": [{"type": "uint256"}], "name": "tokenURI", "outputs": [{"type": "string"}], "stateMutability": "view", "type": "function"}
]`

type NFTContractRequest struct {
	ContractAddress string `json:"contractAddress"`
	ChainName       string `json:"chainName"`
}

type NFTContractResponse struct {
	ContractAddress string            `json:"contractAddress"`
	ChainName       string            `json:"chainName"`
	Details         map[string]string `json:"details"`
}

// New model for storing contract details

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/nftcontract")
	{
		g.Use(paseto.PASETO(true))
		g.POST("", getnftcontractinfo)
	}
}

func getnftcontractinfo(c *gin.Context) {
	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)

	// Check if user has a domain
	var domain models.Domain
	if err := db.Where("id = ?", userId).First(&domain).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			httpo.NewErrorResponse(403, "User has not created a domain yet").SendD(c)
		} else {
			logwrapper.Error("Failed to check user domain", err)
			httpo.NewErrorResponse(500, "Internal server error").SendD(c)
		}
		return
	}

	var request NFTContractRequest
	if err := c.BindJSON(&request); err != nil {
		httpo.NewErrorResponse(400, "Invalid request payload").SendD(c)
		return
	}

	// Check if chain name is "eth"
	if strings.ToLower(request.ChainName) != "eth" {
		httpo.NewErrorResponse(400, "Unsupported chain name. Only 'eth' is supported").SendD(c)
		return
	}

	// Validate contract address
	if !common.IsHexAddress(request.ContractAddress) {
		httpo.NewErrorResponse(400, "Invalid contract address").SendD(c)
		return
	}

	// Connect to Ethereum mainnet
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		logwrapper.Error("Failed to connect to the Ethereum client", err)
		httpo.NewErrorResponse(500, "Failed to connect to the Ethereum network").SendD(c)
		return
	}

	address := common.HexToAddress(request.ContractAddress)

	parsedABI, err := abi.JSON(strings.NewReader(NFTContractABI))
	if err != nil {
		logwrapper.Error("Failed to parse ABI", err)
		httpo.NewErrorResponse(500, "Failed to parse contract ABI").SendD(c)
		return
	}

	details := getNFTContractDetails(client, address, parsedABI)

	// Store contract details in the database
	contractDetails := models.NftSubscription{
		UserID:          userId,
		ContractAddress: request.ContractAddress,
		ChainName:       request.ChainName,
		Name:            details["name"],
		Symbol:          details["symbol"],
		TotalSupply:     details["totalSupply"],
		Owner:           details["owner"],
		TokenURI:        details["tokenURI(1)"],
	}

	if err := db.Create(&contractDetails).Error; err != nil {
		logwrapper.Error("Failed to store contract details", err)
		httpo.NewErrorResponse(500, "Failed to store contract details").SendD(c)
		return
	}

	response := NFTContractResponse{
		ContractAddress: request.ContractAddress,
		ChainName:       request.ChainName,
		Details:         details,
	}

	httpo.NewSuccessResponseP(200, "NFT contract details retrieved and stored successfully", response).SendD(c)
}

func getNFTContractDetails(client *ethclient.Client, address common.Address, parsedABI abi.ABI) map[string]string {
	details := make(map[string]string)
	methods := []string{"name", "symbol", "totalSupply", "owner"}

	var wg sync.WaitGroup
	var mu sync.Mutex

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, method := range methods {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			result, err := callContractMethod(ctx, client, address, parsedABI, m)
			mu.Lock()
			if err != nil {
				details[m] = "N/A"
			} else {
				details[m] = result
			}
			mu.Unlock()
		}(method)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		tokenURIResult, err := callContractMethod(ctx, client, address, parsedABI, "tokenURI", big.NewInt(1))
		mu.Lock()
		if err != nil {
			details["tokenURI(1)"] = "N/A"
		} else {
			details["tokenURI(1)"] = tokenURIResult
		}
		mu.Unlock()
	}()

	wg.Wait()
	return details
}

func callContractMethod(ctx context.Context, client *ethclient.Client, address common.Address, parsedABI abi.ABI, methodName string, args ...interface{}) (string, error) {
	data, err := parsedABI.Pack(methodName, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack data for %s: %v", methodName, err)
	}

	msg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call %s: %v", methodName, err)
	}

	method, exist := parsedABI.Methods[methodName]
	if !exist {
		return "", fmt.Errorf("method %s not found in ABI", methodName)
	}

	output, err := method.Outputs.Unpack(result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result for %s: %v", methodName, err)
	}

	if len(output) == 0 {
		return "", fmt.Errorf("no output for method %s", methodName)
	}

	return fmt.Sprintf("%v", output[0]), nil
}
