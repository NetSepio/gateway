package account

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/NetSepio/gateway/util/httpo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

// For creating account or siging pass without paseto
// For linking pass paseto
func authGoogle(c *gin.Context) {
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
	err = db.Model(&models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User does not exist, so create a new user
			user = models.User{
				Email:  &email,
				UserId: uuid.NewString(),
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

	customClaims := claims.NewWithEmail(user.UserId, user.Email)
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
		Token:  pasetoToken,
		UserId: user.UserId,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
func authGoogleApp(c *gin.Context) {
	db := dbconfig.GetDb()
	var request AppAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	var user models.User
	err = db.Model(&models.User{}).Where("google = ?", request.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// User does not exist, so create a new user
			user = models.User{
				Google: &request.Email,
				UserId: uuid.NewString(),
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

	c.Set(paseto.CTX_USER_ID, user.UserId)

	customClaims := claims.NewWithEmail(user.UserId, user.Email)
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
		Token:  pasetoToken,
		UserId: user.UserId,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", payload).SendD(c)
}
func registerApple(c *gin.Context) {
	// userId := c.GetString(paseto.CTX_USER_ID)
	db := dbconfig.GetDb()
	var request AppAccountRegisterApple
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	// ctx := context.Background()
	// tokenValidationRes, err := verifyIDToken(ctx, request.IdToken)
	// if err != nil {
	// 	httpo.NewErrorResponse(http.StatusBadRequest, "failed to validate ").SendD(c)
	// 	return
	// }

	// response := map[string]interface{}{
	// 	"email": payload.Claims["email"],
	// 	"name":  payload.Claims["name"],
	// }
	// if !tokenValidationRes.Claims["email_verified"].(bool) {
	// 	httpo.NewErrorResponse(http.StatusForbidden, "email not verified").SendD(c)
	// 	return
	// }

	// email := tokenValidationRes.Claims["email"].(string)
	var user models.User
	err = db.Model(&models.User{}).Where("email = ?", request.Email).First(&user).Error
	if err == nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "another user with this email already exist").SendD(c)
		return
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User does not exist, so create a new user
			user = models.User{
				Email:  &request.Email,
				UserId: uuid.NewString(),
				Apple:  &request.AppleId,
			}
			err = db.Model(&models.User{}).Create(&user).Error
			if err != nil {
				logwrapper.Errorf("failed to create user: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
				return
			}
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// // Other error occurred
			// logwrapper.Errorf("failed to retrieve user: %s", err)
			// httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			// return
			if user.UserId != "" {
				err = db.Model(&models.User{}).Where("email = ?", request.Email).Update("email", request.Email).Error
				if err != nil {
					logwrapper.Errorf("failed to update user email: %s", err)
					httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
					return
				}
				httpo.NewSuccessResponse(200, "account Updated successfully").SendD(c)
				return
			} else {
				// Other error occurred
				logwrapper.Errorf("failed to retrieve user: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
				return
			}
		}
	}

	// customClaims := claims.NewWithEmail(user.UserId, user.Email)
	// pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	// if err != nil {
	// 	httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
	// 	logwrapper.Errorf("failed to generate token, error %v", err.Error())
	// 	return
	// }

	// pasetoToken, err := auth.GenerateToken(customClaims, pvKey)
	// if err != nil {
	// 	logwrapper.Errorf("failed to create paseto token: %s", err)
	// 	httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
	// 	return
	// }

	// payload := CreateAccountResponse{
	// 	Token:  pasetoToken,
	// 	UserId: user.UserId,
	// }
	response := map[string]string{
		"data":   "user created scuccessfully",
		"userId": user.UserId,
	}
	httpo.NewSuccessResponseP(200, "Token generated successfully", response).SendD(c)
}

func verifyIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, "699954671747-i1mqa1d7k3nh8bo9arv58pq8j4osl642.apps.googleusercontent.com")
	if err != nil {
		return nil, err
	}
	return payload, nil
}

type UserRequestDetails struct {
	AppleId string `json:"appleId"`
}

func getUserDetails(c *gin.Context) {
	db := dbconfig.GetDb() // Initialize the database connection
	var queryReq UserRequestDetails

	// Bind query parameters from the request
	if err := c.BindJSON(&queryReq); err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Payload is invalid: %s", err)).SendD(c)
		return
	}

	// Validate Email
	if queryReq.AppleId == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "Email is required").SendD(c)
		return
	}

	var user models.User
	// Query the database to find the user by Email
	if err := db.Where("apple_id = ?", queryReq.AppleId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "User not found").SendD(c)
		} else {
			httpo.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Database error: %s", err)).SendD(c)
		}
		return
	}

	// Respond with the user details
	httpo.NewSuccessResponseP(200, "Account fetched successfully", user).SendD(c)
}
