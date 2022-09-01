package claims

import (
	"os"
	"strconv"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
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
	err := db.Model(&models.User{}).Where("wallet_address = ?", c.WalletAddress).First(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func New(walletAddress string) CustomClaims {
	pasetoExpirationInHours, ok := os.LookupEnv("PASETO_EXPIRATION_IN_HOURS")
	pasetoExpirationInHoursInt := time.Duration(24)
	if ok {
		res, err := strconv.Atoi(pasetoExpirationInHours)
		if err != nil {
			logwrapper.Log.Warnf("Failed to parse PASETO_EXPIRATION_IN_HOURS as int : %v", err.Error())
		} else {
			pasetoExpirationInHoursInt = time.Duration(res)
		}
	}
	expiration := time.Now().Add(pasetoExpirationInHoursInt * time.Hour)
	signedBy := envconfig.EnvVars.SIGNED_BY
	return CustomClaims{
		walletAddress,
		signedBy,
		pvx.RegisteredClaims{
			Expiration: &expiration,
		},
	}
}
