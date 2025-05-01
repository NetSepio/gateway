package sdkauthentication

import (
	"encoding/hex"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/models/claims"
	"netsepio-gateway-v1.1/utils/auth"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

var (
	accessKey = os.Getenv("ACCESS_KEY")
	// secretKeyHex = os.Getenv("PASETO_SECRETKEYHEX") // Replace with your actual private key
)

// var secretKey paseto.V4AsymmetricSecretKey

// func init() {
// 	var err error
// 	secretKey, err = paseto.NewV4AsymmetricSecretKeyFromHex(secretKeyHex)
// 	if err != nil {
// 		log.Fatalf("Failed to create secret key: %v", err)
// 	}
// }

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/sdkauthentication")
	{
		g.GET("/generate-token", generateToken)
	}
}

func generateToken(c *gin.Context) {
	db := database.GetDb()

	providedAccessKey := c.Query("access_key")
	walletAddress := c.Query("wallet_address")

	if providedAccessKey != accessKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access key"})
		return
	}

	var userData = models.User{}

	err := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).First(&userData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Record Not Found").SendD(c)
		logwrapper.Errorf("Record Not Found, wallet_address:  %v", walletAddress)
		return
	}
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Record Not Found").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}

	customClaims := claims.NewWithWallet(userData.UserId, &walletAddress)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}
	pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}

	// Create a new PASETO token
	// token := paseto.NewToken()
	// token.SetIssuedAt(time.Now())
	// token.SetExpiration(time.Now().Add(24 * time.Hour))
	// token.SetString("wallet_address", walletAddress)
	// token.SetString("access_key", providedAccessKey)

	// // Sign the token
	// signed := token.V4Sign(secretKey, nil)

	c.JSON(http.StatusOK, gin.H{"token": pasetoToken})
}
