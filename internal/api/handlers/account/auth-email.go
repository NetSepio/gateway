package account

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/models/claims"
	"netsepio-gateway-v1.1/utils/auth"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func GenerateAuthId(c *gin.Context) {
	db := database.GetDb()
	// parse request
	var request GenerateAuthIdRequest
	err := c.BindJSON(&request)
	if err != nil {

		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	// delete all auth tokens which are expired
	intervalMin := fmt.Sprintf("%.0f minutes", load.Cfg.MAGIC_LINK_EXPIRATION.Minutes())
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

	auth := smtp.PlainAuth("", "apikey", load.Cfg.SMTP_PASSWORD, smtpHost)
	appSubDomain := "app"
	if load.Cfg.NETWORK == "testnet" {
		appSubDomain = "dev"
	}

	appLink := fmt.Sprintf("https://%s.netsepio.com/magiclink?token=%s", appSubDomain, authCode)
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
	db := database.GetDb()
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
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

	intervalMin := fmt.Sprintf("%.0f minutes", load.Cfg.MAGIC_LINK_EXPIRATION.Minutes())

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
	userId := c.GetString(paseto.CTX_USER_ID)
	fmt.Printf("userId: %s\n", userId)
	// don't create user if paseto exist
	var user models.User
	err = db.Model(&models.User{}).Where("email = ?", emailAuth.Email).First(&user).Error
	if err == nil {
		if userId != "" {
			// return error stating that user with email exist so it cannot be linked
			httpo.NewErrorResponse(http.StatusConflict, "User with email already exist").SendD(c)
			return
		}
		// user exist, so generate paseto for that user id
		customClaims := claims.NewWithEmail(user.UserId, user.Email)
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

	if userId != "" {
		// update user with that email
		if err = db.Model(&models.User{}).Where("user_id = ?", userId).Update("email", emailAuth.Email).Error; err != nil {
			logwrapper.Errorf("failed to update user: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			c.Abort()
			return
		}
		//delete all records for that email in email auth
		err = db.Model(&models.EmailAuth{}).Where("email = ?", emailAuth.Email).Delete(&models.EmailAuth{}).Error
		if err != nil {
			logwrapper.Errorf("failed to delete email auth: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
		httpo.NewSuccessResponse(200, "User linked successfully").SendD(c)
		return
	}
	newUserId := uuid.NewString()

	// create user with that email
	if err = db.Create(&models.User{Email: &emailAuth.Email, UserId: newUserId}).Error; err != nil {
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
