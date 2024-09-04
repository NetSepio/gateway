package dvpnnft

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"regexp"

	contract "github.com/NetSepio/gateway/api/v1/dvpnnft/contract"
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
	Message         string `json:"message,omitempty"`
}
type APTRequestPayload struct {
	WalletAddress string `json:"wallet_address"`
	EmailID       string `json:"email_id"`
}

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/dvpnnft")
	{
		g.POST("/chain=evm", handleMintNFTEVM)
		g.POST("/chain=apt", handleMintNFTAPT)
	}
}
func handleMintNFTEVM(c *gin.Context) {
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
	if err := storeNFTRecord("evm", payload.WalletAddress, "", tx.Hash().Hex()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing transaction hash in the database"})
		return
	}
	c.JSON(http.StatusOK, ResponsePayload{
		TransactionHash: tx.Hash().Hex(),
	})
}

// func storeTransactionHash(walletAddress, transactionHash string) error {
// 	db := dbconfig.GetDb()
// 	nft := &models.DVPNNFTRecord{
// 		WalletAddress:   walletAddress,
// 		TransactionHash: transactionHash,
// 	}
// 	return db.Create(nft).Error
// }

//-------------------------aptos -------------------------------------------

func handleMintNFTAPT(c *gin.Context) {
	var payload APTRequestPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	fmt.Printf("Received Aptos address: %s\n", payload.WalletAddress)
	fmt.Printf("Is valid Aptos address: %v\n", isValidAptosAddress(payload.WalletAddress))

	// Check if the wallet address and email are valid and don't exist in the database
	if !isValidAptosAddress(payload.WalletAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Aptos wallet address"})
		return
	}

	if !isValidEmail(payload.EmailID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	if exists, err := checkWalletAddressExists(payload.WalletAddress, "apt"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	} else if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Wallet address already exists"})
		return
	}

	if exists, err := checkEmailExists(payload.EmailID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	} else if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email address already exists"})
		return
	}

	// TODO: Implement Aptos minting logic here

	// For now, we'll just store the record without a transaction hash
	if err := storeNFTRecord("apt", payload.WalletAddress, payload.EmailID, ""); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing data"})
		return
	}

	c.JSON(http.StatusOK, ResponsePayload{
		Message: "Aptos NFT minting record created",
	})
}

func isValidMantaAddress(address string) bool {
	mantaAddressRegex := `^0x[0-9a-fA-F]{40}$`
	match, _ := regexp.MatchString(mantaAddressRegex, address)
	return match
}

func isValidAptosAddress(address string) bool {
	aptosAddressRegex := `^0x[0-9a-fA-F]{64}$`
	match, _ := regexp.MatchString(aptosAddressRegex, address)
	return match
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func checkWalletAddressExists(walletAddress, chain string) (bool, error) {
	db := dbconfig.GetDb()
	var count int64
	result := db.Model(&models.DVPNNFTRecord{}).Where("wallet_address = ? AND chain = ?", walletAddress, chain).Count(&count)
	return count > 0, result.Error
}

func checkEmailExists(email string) (bool, error) {
	db := dbconfig.GetDb()
	var count int64
	result := db.Model(&models.DVPNNFTRecord{}).Where("email_id = ?", email).Count(&count)
	return count > 0, result.Error
}

func storeNFTRecord(chain, walletAddress, emailID, transactionHash string) error {
	db := dbconfig.GetDb()
	nft := &models.DVPNNFTRecord{
		Chain:           chain,
		WalletAddress:   walletAddress,
		EmailID:         emailID,
		TransactionHash: transactionHash,
	}
	return db.Create(nft).Error
}
