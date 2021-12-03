package claims

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	WalletAddress string `json:"walletAddress"`
	jwt.RegisteredClaims
}

func New(walletAddress string) CustomClaims {
	return CustomClaims{
		walletAddress,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
}
