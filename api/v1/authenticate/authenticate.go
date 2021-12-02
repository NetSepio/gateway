package authenticate

import (
	"fmt"
	"net/http"
	"netsepio-api/db"
	"netsepio-api/models"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {

	//TODO remove flow id if 200
	var req AuthenticateRequest
	c.BindJSON(&req)

	// Append userId to the message
	message := req.FlowId + "m"
	fmt.Println("flowid", req.FlowId)
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(req.Signature)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	signatureInBytes[64] -= 27
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)
	var user models.User
	res := db.Db.Model(&models.User{}).Where("? = ANY (flow_id)", req.FlowId).First(&user)
	if res.RecordNotFound() {
		c.String(http.StatusNotFound, "FlowId not found")
		return
	}
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(user.WalletAddress, "==", walletAddress.String())
	if user.WalletAddress == walletAddress.String() {
		jwtToken, err := generateToken(user.WalletAddress)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error Occured")
		}
		c.JSON(http.StatusOK, map[string]string{
			"token": jwtToken,
		})
	} else {
		c.String(http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}

func generateToken(walletAddress string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, customClaims{
		walletAddress,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("PASETO_PRIVATE_KEY")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
