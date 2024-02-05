package account

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vk-rv/pvx"
)

func GenerateAuthId(c *gin.Context) {
	db := dbconfig.GetDb()

	// parse request
	var request GenerateAuthIdRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	//delete all records for that email in email auth
	err = db.Model(&models.EmailAuth{}).Where("email = ?", request.Email).Delete(&models.EmailAuth{}).Error
	if err != nil {
		logwrapper.Errorf("failed to delete email auth: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	// generate auth id
	authId := uuid.NewString()
	// save auth id in db
	emailAuth := models.EmailAuth{
		Id:    authId,
		Email: request.Email,
	}
	if err := db.Create(&emailAuth).Error; err != nil {
		logwrapper.Errorf("failed to save auth id in db: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	//generate paseto with authid
	pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to decode priv key, %s", err)
		return
	}

	customClaims := claims.NewAuthClaim(authId)
	pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		logwrapper.Errorf("failed to create paseto token: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	smtpHost := "localhost"
	smtpPort := "465"
	smtpFrom := "mail@netsepio.com"
	auth := smtp.PlainAuth("NetSepio", smtpFrom, envconfig.EnvVars.SMTP_PASSWORD, smtpHost)
	appSubDomain := "app"
	if envconfig.EnvVars.NETWORK == "testnet" {
		appSubDomain = "dev"
	}
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, []string{request.Email}, []byte(fmt.Sprintf("Subject: Magic Link\n\nhttps://%s.netsepio.com/magiclink?token=%s", appSubDomain, pasetoToken)))
	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	httpo.NewSuccessResponse(200, "Magic link send").SendD(c)
}

// API controller which takes paseto from request, then validates it using claims valid
func PasetoFromMagicLink(c *gin.Context) {
	db := dbconfig.GetDb()
	pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to decode priv key, %s", err)
		return
	}
	var request PasetoFromMagicLinkRequest
	err = c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}
	ppv4 := pvx.NewPV4Public()
	pubKey := ed25519.PrivateKey(pvKey).Public().(ed25519.PublicKey)
	asymPK := pvx.NewAsymmetricPublicKey(pubKey, pvx.Version4)
	var cc claims.AuthClaim
	err = ppv4.
		Verify(request.Token, asymPK).
		ScanClaims(&cc)

	if err != nil {
		var validationErr *pvx.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.HasExpiredErr() {
				logwrapper.Errorf("failed to scan claims for paseto token: %s", err)
				httpo.NewErrorResponse(httpo.TokenExpired, "token expired").Send(c, http.StatusUnauthorized)
				c.Abort()
				return
			}
			logwrapper.Errorf("failed to scan claims for paseto token: %s", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		logwrapper.Errorf("failed to scan claims for paseto token: %s", err)
		httpo.NewErrorResponse(http.StatusUnauthorized, "invalid token").SendD(c)
		c.Abort()
		return

	} else {
		if err := cc.Valid(); err != nil {
			logwrapper.Errorf("failed to validate paseto token: %s", err)
			httpo.NewErrorResponse(http.StatusUnauthorized, "invalid token").SendD(c)
			c.Abort()
			return
		}

		// get email from db for authId
		var emailAuth models.EmailAuth
		err := db.Model(&models.EmailAuth{}).Where("id = ?", cc.AuthId).First(&emailAuth).Error
		if err != nil {
			logwrapper.Errorf("failed to get email from db: %s", err)
			httpo.NewErrorResponse(http.StatusUnauthorized, "invalid token").SendD(c)
			c.Abort()
			return
		}

		// don't create user if paseto exist
		var user models.User
		err = db.Model(&models.User{}).Where("email_id = ?", emailAuth.Email).First(&models.User{}).Error
		if err == nil {
			// user exist, so generate paseto for that user id
			customClaims := claims.NewWithEmail(user.UserId, user.EmailId)
			pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
			if err != nil {
				logwrapper.Errorf("failed to create paseto token: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
				return
			}
			// send paseto in response
			payload := PasetoFromMagicLinkResponse{
				Token: pasetoToken,
			}
			//delete all records for that email in email auth
			err = db.Model(&models.EmailAuth{}).Where("email = ?", emailAuth.Email).Delete(&models.EmailAuth{}).Error
			if err != nil {
				logwrapper.Errorf("failed to delete email auth: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
				return
			}
			httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
			return
		}
		newUserId := uuid.NewString()

		// create user with that email
		if err = db.Create(&models.User{EmailId: &emailAuth.Email, UserId: newUserId}).Error; err != nil {
			logwrapper.Errorf("failed to create user: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			c.Abort()
			return
		}

		// create paseto for that userId
		customClaims := claims.NewWithEmail(newUserId, &emailAuth.Email)
		pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
		if err != nil {
			logwrapper.Errorf("failed to create paseto token: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
		// send paseto in response
		payload := PasetoFromMagicLinkResponse{
			Token: pasetoToken,
		}
		//delete all records for that email in email auth
		err = db.Model(&models.EmailAuth{}).Where("email = ?", emailAuth.Email).Delete(&models.EmailAuth{}).Error
		if err != nil {
			logwrapper.Errorf("failed to delete email auth: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
		httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
		return
	}

}
