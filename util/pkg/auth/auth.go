package auth

import (
	"github.com/vk-rv/pvx"
)

func GenerateToken(claims pvx.Claims, privateKey string) (string, error) {
	asymSK := pvx.NewAsymmetricSecretKey([]byte(privateKey), pvx.Version4)
	ppv4 := pvx.NewPV4Public()
	tokenString, err := ppv4.Sign(asymSK, claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
