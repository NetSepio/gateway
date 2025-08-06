package paseto

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/vk-rv/pvx"

	"github.com/gin-gonic/gin"
)

var CTX_WALLET_ADDRES = "WALLET_ADDRESS"
var CTX_USER_ID = "USER_ID"
var CTX_ORGANISATION_ID = "ORGANISATION_ID"
var CTX_ORGANISATION_IP = "ORGANISATION_IP"
var CTX_ORGANISATION_NAME = "ORGANISATION_NAME"
var CTX_ORG_APP_ID = "ORGANISATION_APP_ID"
var CTX_VERIFIED = "VERIFY"

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(authOptional bool) func(*gin.Context) {
	return func(c *gin.Context) {
		var headers GenericAuthHeaders
		err := c.BindHeader(&headers)
		if err != nil {
			err = fmt.Errorf("failed to bind header, %s", err)
			logValidationFailed(headers.Authorization, err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if headers.Authorization == "" {
			if authOptional {
				c.Next()
				return
			}
			logValidationFailed(headers.Authorization, ErrAuthHeaderMissing)
			httpo.NewErrorResponse(http.StatusBadRequest, ErrAuthHeaderMissing.Error()).SendD(c)
			c.Abort()
			return
		} else if !strings.HasPrefix(headers.Authorization, "Bearer ") {
			err := errors.New("authorization header must have Bearer prefix")
			logValidationFailed(headers.Authorization, err)
			httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
			c.Abort()
			return
		}

		pasetoToken := strings.TrimPrefix(headers.Authorization, "Bearer ")
		ppv4 := pvx.NewPV4Public()

		k, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			err = fmt.Errorf("failed to decode priv key, %s", err)
			logValidationFailed(headers.Authorization, err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		pubKey := ed25519.PrivateKey(k).Public().(ed25519.PublicKey)
		asymPK := pvx.NewAsymmetricPublicKey(pubKey, pvx.Version4)

		var cc claims.CustomClaims
		err = ppv4.Verify(pasetoToken, asymPK).ScanClaims(&cc)
		if err == nil && cc.Valid() == nil {
			// Valid CustomClaims
			db := database.GetDb()
			var userFetch models.User
			err := db.Model(&models.User{}).Where("user_id = ?", strings.ToLower(cc.UserId)).First(&userFetch).Error
			if err != nil {
				err = fmt.Errorf("failed to get wallet address, %s", err)
				logwrapper.Log.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if userFetch.WalletAddress != nil {
				c.Set(CTX_WALLET_ADDRES, *userFetch.WalletAddress)
			}
			c.Set(CTX_USER_ID, userFetch.UserId)
			c.Next()
			return
		}

		// Try parsing as CustomClaimsForOrganisation
		var occ claims.CustomClaimsForOrganisation
		err = ppv4.Verify(pasetoToken, asymPK).ScanClaims(&occ)
		if err == nil && occ.Valid() == nil {
			// Set context for organisation
			c.Set(CTX_ORGANISATION_ID, occ.OrganisationId)
			if occ.OrganisationName != nil {
				c.Set(CTX_ORGANISATION_NAME, occ.OrganisationName)
			}
			if occ.IpAddress != nil {
				c.Set(CTX_ORGANISATION_IP, *occ.IpAddress)
			}
			c.Next()
			return
		}

		// Try parsing as CustomClaimsForOrganisationApp
		var occApp claims.CustomClaimsForOrganisationApp
		err = ppv4.Verify(pasetoToken, asymPK).ScanClaims(&occApp)
		if err == nil && occApp.Valid() == nil {
			// Set context for organisation app
			c.Set(CTX_ORGANISATION_ID, occApp.AppOrganisationId)
			c.Set(CTX_ORG_APP_ID, occApp.AppOrganisationId)
			// if occApp.OrganisationName != nil {
			// 	c.Set(CTX_ORGANISATION_NAME, occApp.OrganisationName)
			// }
			// if occApp.IpAddress != nil {
			// 	c.Set(CTX_ORGANISATION_IP, *occApp.IpAddress)
			// }
			c.Next()
			return
		}

		// Handle validation errors
		var validationErr *pvx.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.HasExpiredErr() {
				err = fmt.Errorf("token expired")
				logValidationFailed(headers.Authorization, err)
				httpo.NewErrorResponse(httpo.TokenExpired, "token expired").Send(c, http.StatusUnauthorized)
				c.Abort()
				return
			}
		}
		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func logValidationFailed(token string, err error) {
	logwrapper.Warnf("validation failed with token %v and error: %v", token, err)
}
