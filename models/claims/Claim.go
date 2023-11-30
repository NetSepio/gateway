package claims

import (
	"strings"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/vk-rv/pvx"
)

type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	SignedBy      string `json:"signedBy"`
	pvx.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	db := dbconfig.GetDb()
	if err := c.RegisteredClaims.Valid(); err != nil {
		return err
	}
	err := db.Model(&models.User{}).Where("wallet_address = ?", strings.ToLower(c.WalletAddress)).First(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func New(walletAddress string) CustomClaims {
	pasetoExpirationInHours := envconfig.EnvVars.PASETO_EXPIRATION
	expiration := time.Now().Add(pasetoExpirationInHours)
	signedBy := envconfig.EnvVars.PASETO_SIGNED_BY
	return CustomClaims{
		walletAddress,
		signedBy,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}
