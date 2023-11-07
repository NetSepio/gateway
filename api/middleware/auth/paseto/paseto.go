package paseto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/vk-rv/pvx"
	"gorm.io/gorm"

	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		err = fmt.Errorf("failed to bind header, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(headers.Authorization, ErrAuthHeaderMissing)
		httpo.NewErrorResponse(http.StatusBadRequest, ErrAuthHeaderMissing.Error()).SendD(c)
		c.Abort()
		return
	}
	pasetoToken := headers.Authorization
	pv4 := pvx.NewPV4Local()
	k := envconfig.EnvVars.PASETO_PRIVATE_KEY
	symK := pvx.NewSymmetricKey([]byte(k), pvx.Version4)
	var cc claims.CustomClaims
	err = pv4.
		Decrypt(pasetoToken, symK).
		ScanClaims(&cc)
	if err != nil {
		var validationErr *pvx.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.HasExpiredErr() {
				err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
				logValidationFailed(headers.Authorization, err)
				httpo.NewErrorResponse(http.StatusUnauthorized, "token expired").Send(c, httpo.TokenExpired)
				c.Abort()
				return
			}

		}
		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		if err := cc.Valid(); err != nil {
			logValidationFailed(headers.Authorization, err)
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				err = fmt.Errorf("failed to validate claim, %s", err)
				logwrapper.Log.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		} else {
			c.Set("walletAddress", cc.WalletAddress)
			c.Next()
		}
	}
}

func logValidationFailed(token string, err error) {
	logwrapper.Warnf("validation failed with token %v and error: %v", token, err)
}
