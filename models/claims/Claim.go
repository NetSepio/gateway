package claims

import (
	"os"
	"strconv"
	"time"

	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	SignedBy      string `json:"signedBy"`
	jwt.RegisteredClaims
}

func New(walletAddress string) CustomClaims {
	jwtExpirationInHours, ok := os.LookupEnv("JWT_EXPIRATION_IN_HOURS")
	jwtExpirationInHoursInt := time.Duration(24)
	if ok {
		res, err := strconv.Atoi(jwtExpirationInHours)
		if err != nil {
			logwrapper.Log.Warnf("Failed to parse JWT_EXPIRATION_IN_HOURS as int : %v", err.Error())
		} else {
			jwtExpirationInHoursInt = time.Duration(res)
		}
	}
	signedBy := envutil.MustGetEnv("SIGNED_BY")
	return CustomClaims{
		walletAddress,
		signedBy,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationInHoursInt * time.Hour)),
		},
	}
}
