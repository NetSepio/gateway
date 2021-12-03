package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims jwt.Claims, privateKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(privateKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
