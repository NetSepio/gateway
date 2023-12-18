package account

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

func create(c *gin.Context) {
	db := dbconfig.GetDb()
	var request CreateAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	tokenValidationRes, err := idtoken.Validate(context.Background(), request.IdToken, envconfig.EnvVars.GOOGLE_AUDIENCE)
	if err != nil {
		logwrapper.Errorf("failed to validate id token: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	if !tokenValidationRes.Claims["email_verified"].(bool) {
		httpo.NewErrorResponse(http.StatusForbidden, "email not verified").SendD(c)
		return
	}

	email := tokenValidationRes.Claims["email"].(string)
	var user models.User
	err = db.Model(&models.User{}).Where("email_id = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User does not exist, so create a new user
			user = models.User{
				EmailId: email,
				UserId:  uuid.NewString(),
			}
			err = db.Model(&models.User{}).Create(&user).Error
			if err != nil {
				logwrapper.Errorf("failed to create user: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
				return
			}
		} else {
			// Other error occurred
			logwrapper.Errorf("failed to retrieve user: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
	}

	customClaims := claims.New(user.UserId, &user.EmailId)
	pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to generate token, error %v", err.Error())
		return
	}

	pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		logwrapper.Errorf("failed to create paseto token: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	payload := CreateAccountResponse{
		Token: pasetoToken,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
