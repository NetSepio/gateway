package account

import (
	"context"
	"encoding/hex"
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
)

func create(c *gin.Context) {
	db := dbconfig.GetDb()
	var request CreateAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}
	tokenValidationRes, err := idtoken.Validate(context.Background(), request.IdToken, envconfig.EnvVars.GOOGLE_AUDIENCE)
	if err != nil {
		panic(err)
	}

	if !tokenValidationRes.Claims["email_verified"].(bool) {
		httpo.NewErrorResponse(http.StatusForbidden, "email not verified").SendD(c)
		return
	}
	// create user
	user := models.User{
		EmailId: tokenValidationRes.Claims["email"].(string),
		UserId:  uuid.NewString(),
	}
	err = db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		logwrapper.Errorf("failed to create user: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}
	customClaims := claims.New(user.UserId, nil)
	// create paseto
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

	//send paseto as success response
	payload := CreateAccountResponse{
		Token: pasetoToken,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
