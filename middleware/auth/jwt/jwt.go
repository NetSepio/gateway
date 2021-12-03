package jwt

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func JWT(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		c.String(http.StatusBadRequest, "Authorization header is required")
		return
	}
	jwtToken := headers.Authorization[7:]
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		jwtPrivateKet := []byte(os.Getenv("JWT_PRIVATE_KEY"))
		return jwtPrivateKet, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("walletAddress", claims["walletAddress"])
		c.Next()
	} else {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusForbidden)
	}
}
