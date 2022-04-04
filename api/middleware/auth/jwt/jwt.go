package jwt

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"

	"github.com/NetSepio/gateway/util/pkg/envutil"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func JWT(c *gin.Context) {
	db := dbconfig.GetDb()
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(headers.Authorization, err)
		httphelper.ErrResponse(c, http.StatusBadRequest, ErrAuthHeaderMissing.Error())
		c.Abort()
		return
	}
	jwtToken := headers.Authorization[7:]
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		jwtPrivateKet := []byte(envutil.MustGetEnv("JWT_PRIVATE_KEY"))
		return jwtPrivateKet, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		walletAddress := claims["walletAddress"]

		err := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress.(string)).First(&models.User{}).Error
		if err != nil {
			logValidationFailed(headers.Authorization, err)
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
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func logValidationFailed(token string, err error) {
	logwrapper.Warnf("validation failed with token %v and error: %v", token, err)
}
