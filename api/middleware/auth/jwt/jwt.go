package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/db"
	"github.com/TheLazarusNetwork/marketplace-engine/models"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func JWT(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		httphelper.ErrResponse(c, http.StatusBadRequest, ErrAuthHeaderMissing.Error())
		c.Abort()
		return
	}
	jwtToken := headers.Authorization[7:]
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		jwtPrivateKet := []byte(os.Getenv("JWT_PRIVATE_KEY"))
		return jwtPrivateKet, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		walletAddress := claims["walletAddress"]

		err := db.Db.Model(&models.User{}).Where("wallet_address = ?", walletAddress.(string)).First(&models.User{}).Error
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				c.AbortWithStatus(http.StatusForbidden)
			} else {
				logwrapper.Log.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		} else {
			c.Set("walletAddress", walletAddress)
			c.Next()
		}
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
