package auth

import (
	"github.com/vk-rv/pvx"
)

func GenerateToken(claims pvx.Claims, privateKey string) (string, error) {

	symK := pvx.NewSymmetricKey([]byte(privateKey), pvx.Version4)
	pv4 := pvx.NewPV4Local()
	tokenString, err := pv4.Encrypt(symK, claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
