package account

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// delete all auth tokens which are expired
	intervalMin := fmt.Sprintf("%.0f minutes", envconfig.EnvVars.MAGIC_LINK_EXPIRATION.Minutes())
	whereQuery := fmt.Sprintf("created_at >= NOW() - INTERVAL '%s minutes'", intervalMin)
	err = db.Model(&models.EmailAuth{}).Where(whereQuery).Delete(&models.EmailAuth{}).Error
	if err != nil {
		logwrapper.Errorf("failed to delete expired email auth: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
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
	authCode := uuid.NewString()
	authCode = strings.ReplaceAll(authCode, "-", "")
	authCode = authCode[:6]
	// save auth id in db
	emailAuth := models.EmailAuth{
		Id:       authId,
		Email:    request.Email,
		AuthCode: authCode,
	}
	if err := db.Create(&emailAuth).Error; err != nil {
		logwrapper.Errorf("failed to save auth id in db: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	smtpHost := "smtp.sendgrid.net"
	smtpPort := 587
	smtpFrom := "support@netsepio.com"

	auth := smtp.PlainAuth("", "apikey", envconfig.EnvVars.SMTP_PASSWORD, smtpHost)
	appSubDomain := "app"
	if envconfig.EnvVars.NETWORK == "testnet" {
		appSubDomain = "dev"
	}

	appLink := fmt.Sprintf("https://%s.netsepio.com/maginlink?token=%s", appSubDomain, authCode)
	// Create body for sending auth token and link with auth token
	body := fmt.Sprintf(`Dear user,<br>Login was requested for this email, click this link to login %s<br>Alternatively you can also enter this code into platform %s<br>Don’t share this code and link with anyone.`, appLink, authCode)
	// The msg parameter should be an RFC 822-style email with headers first, a blank line, and then the message body. The lines of msg should be CRLF terminated. The msg headers should usually include fields such as "From", "To", "Subject", and "Cc".
	// construct message as per rpc for mail including things needed like From To Subject
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s", smtpFrom, request.Email, "Login to NetSepio", body)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, smtpFrom, []string{request.Email}, []byte(msg))
	// handling the errors
	if err != nil {
		logwrapper.Errorf("failed to send email: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "Magiclink send").SendD(c)
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

	intervalMin := fmt.Sprintf("%.0f minutes", envconfig.EnvVars.MAGIC_LINK_EXPIRATION.Minutes())

	// get email from db for authId
	var emailAuth models.EmailAuth

	whereQuery := fmt.Sprintf("email = ? AND auth_code = ? AND created_at >= NOW() - INTERVAL '%s minutes'", intervalMin)
	// query with created at > 5 minutes
	err = db.Model(&models.EmailAuth{}).
		Where(whereQuery,
			request.EmailId, request.Code).First(&emailAuth).Error
	if err != nil {
		logwrapper.Errorf("failed to get email from db: %s", err)
		httpo.NewErrorResponse(http.StatusUnauthorized, "invalid code").SendD(c)
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
}
