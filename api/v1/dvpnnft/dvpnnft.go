package dvpnnft

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"regexp"

	contract "github.com/NetSepio/gateway/api/v1/dvpnnft/contract" // Replace with the actual path to your contract bindings
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	migrate "github.com/NetSepio/gateway/models/Migrate"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestPayload struct {
	WalletAddress string `json:"wallet_address"`
}

type ResponsePayload struct {
	TransactionHash string `json:"transaction_hash"`
}

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/dvpnnft")
	{
		g.POST("", handleMintNFT)
	}
}

func handleMintNFT(c *gin.Context) {
	var payload RequestPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// execute db process check the wallet address exist

	db := dbconfig.GetDb()

	// Check if the wallet address exists
	result := db.Where("wallet_address = ?", payload.WalletAddress).First(&migrate.DVPNNFTRecord{})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("Wallet address does not exist")
		} else {
			fmt.Println("Error occurred:", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": result.Error.Error(),
			})
			return
		}
	} else {
		fmt.Println("Wallet address exists:", payload.WalletAddress)
		c.JSON(http.StatusFound, gin.H{
			"warning":        "this wallet address is already minted",
			"wallet address": payload.WalletAddress,
		})
		return
	}

	// Check if the wallet address is a valid Manta address
	if !isValidMantaAddress(payload.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Manta wallet address"})
		return
	}
	// Connect to the Ethereum client
	client, err := ethclient.Dial("https://pacific-rpc.manta.network/http")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
		return
	}

	// Load the private key from the environment
	privateKey, err := crypto.HexToECDSA(envconfig.EnvVars.PRIVATE_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading private key"})
		return
	}

	// Get the public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error casting public key to ECDSA"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(c, fromAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting nonce"})
		return
	}

	gasPrice, err := client.SuggestGasPrice(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting gas price"})
		return
	}

	chainID, err := client.NetworkID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting chain ID"})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating transactor"})
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // Adjust as needed
	auth.GasPrice = gasPrice

	// Contract address
	contractAddress := common.HexToAddress(envconfig.EnvVars.CONTRACT_ADDRESS)
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating contract instance"})
		return
	}

	// Call the mint function
	tx, err := instance.DelegateMint(auth, common.HexToAddress(payload.WalletAddress))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling Mint function"})
		return
	}

	// Store the transaction hash in the database
	if err := storeTransactionHash(payload.WalletAddress, tx.Hash().Hex()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing transaction hash in the database"})
		return
	}

	c.JSON(http.StatusOK, ResponsePayload{
		TransactionHash: tx.Hash().Hex(),
	})
}

func storeTransactionHash(walletAddress, transactionHash string) error {
	db := dbconfig.GetDb()
	nft := &models.DVPNNFTRecord{
		WalletAddress:   walletAddress,
		TransactionHash: transactionHash,
	}
	return db.Create(nft).Error
}
func isValidMantaAddress(address string) bool {
	// Manta address format: 0x[0-9a-fA-F]{40}
	mantaAddressRegex := `^0x[0-9a-fA-F]{40}$`
	match, err := regexp.MatchString(mantaAddressRegex, address)
	if err != nil {
		return false
	}
	return match
}

// package dvpnnft

// import (
// 	"crypto/ecdsa"
// 	"math/big"
// 	"net/http"
// 	"os"
// 	"regexp"

// 	contract "github.com/NetSepio/gateway/api/v1/dvpnnft/contract" // Replace with the actual path to your contract bindings
// 	"github.com/NetSepio/gateway/config/dbconfig"
// 	"github.com/NetSepio/gateway/models"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// type RequestPayload struct {
// 	WalletAddress string `json:"wallet_address"`
// }

// type ResponsePayload struct {
// 	TransactionHash string `json:"transaction_hash"`
// }

// func ApplyRoutes(r *gin.RouterGroup) {
// 	g := r.Group("/dvpnnft")
// 	{
// 		g.POST("", handleMintNFT)
// 	}
// }

// func handleMintNFT(c *gin.Context) {
// 	var payload RequestPayload
// 	if err := c.BindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	// Load environment variables from the .env file
// 	if err := godotenv.Load(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading .env file"})
// 		return
// 	}

// 	// Check if the wallet address is a valid Manta address
// 	if !isValidMantaAddress(payload.WalletAddress) {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Manta wallet address"})
// 		return
// 	}

// 	// Connect to the Ethereum client
// 	client, err := ethclient.Dial("https://rpc-amoy.polygon.technology/")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
// 		return
// 	}

// 	// Load the private key from the environment
// 	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading private key"})
// 		return
// 	}

// 	// Get the public key and address
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error casting public key to ECDSA"})
// 		return
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := client.PendingNonceAt(c, fromAddress)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting nonce"})
// 		return
// 	}

// 	gasPrice, err := client.SuggestGasPrice(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting gas price"})
// 		return
// 	}

// 	chainID, err := client.NetworkID(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting chain ID"})
// 		return
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating transactor"})
// 		return
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // in wei
// 	auth.GasLimit = uint64(300000) // Adjust as needed
// 	auth.GasPrice = gasPrice

// 	// Contract address
// 	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
// 	instance, err := contract.NewContract(contractAddress, client)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating contract instance"})
// 		return
// 	}

// 	// Call the mint function
// 	tx, err := instance.DelegateMint(auth, common.HexToAddress(payload.WalletAddress))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling Mint function"})
// 		return
// 	}

// 	// Store the transaction hash in the database
// 	if err := storeTransactionHash(payload.WalletAddress, tx.Hash().Hex()); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing transaction hash in the database"})
// 		return
// 	}

// 	// Wait for the transaction to be mined
// 	_, err = client.TransactionReceipt(c, tx.Hash())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error waiting for transaction to be mined"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, ResponsePayload{
// 		TransactionHash: tx.Hash().Hex(),
// 	})
// }

// func storeTransactionHash(walletAddress, transactionHash string) error {
// 	db := dbconfig.GetDb()
// 	nft := &models.DVPNNFTRecord{
// 		WalletAddress:   walletAddress,
// 		TransactionHash: transactionHash,
// 	}
// 	return db.Create(nft).Error
// }

// func isValidMantaAddress(address string) bool {
// 	// Manta address format: 0x[0-9a-fA-F]{40}
// 	mantaAddressRegex := `^0x[0-9a-fA-F]{40}$`
// 	match, err := regexp.MatchString(mantaAddressRegex, address)
// 	if err != nil {
// 		return false
// 	}
// 	return match
// }
