package claims

import (
	"netsepio-api/util/pkg/logwrapper"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
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
	return CustomClaims{
		walletAddress,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationInHoursInt * time.Hour)),
		},
	}
}
