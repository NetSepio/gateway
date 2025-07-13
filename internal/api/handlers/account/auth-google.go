package account

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/internal/api/handlers/referral"
	useractivity "github.com/NetSepio/gateway/internal/api/handlers/userActivity"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/utils/actions"
	"github.com/NetSepio/gateway/utils/auth"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/NetSepio/gateway/utils/module"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

// For creating account or siging pass without paseto
// For linking pass paseto
func authGoogle(c *gin.Context) {
	db := database.GetDb()
	var request CreateAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	tokenValidationRes, err := idtoken.Validate(context.Background(), request.IdToken, load.Cfg.GOOGLE_AUDIENCE)
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
			metadata := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
			go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Created, Metadata: &metadata})
		} else {
			// Other error occurred
			logwrapper.Errorf("failed to retrieve user: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
	}

	customClaims := claims.NewWithEmail(user.UserId, user.Email)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
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
	db := database.GetDb()
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
			meta := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
			go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Created, Metadata: &meta})

		} else {
			// Other error occurred
			logwrapper.Errorf("failed to retrieve user: %s", err)
			httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
			return
		}
	}

	c.Set(paseto.CTX_USER_ID, user.UserId)

	customClaims := claims.NewWithEmail(user.UserId, user.Email)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
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
func allAuthApp(c *gin.Context) {
	db := database.GetDb()
	var request AuthAppAccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	if conditions := request.AuthType == AUTH_GOOGLE_APP && request.Email == ""; conditions {
		if request.AppleID != "" {
			httpo.NewErrorResponse(http.StatusBadRequest, "Apple ID is not required for Google App Auth").SendD(c)
			return
		} else {
			httpo.NewErrorResponse(http.StatusBadRequest, "Email is required for Google App Auth").SendD(c)
			return
		}
	} else if conditions := request.AuthType == AUTH_APPLE_APP && request.AppleID == ""; conditions {
		httpo.NewErrorResponse(http.StatusBadRequest, "Apple Id and Apple Email is required for apple authentication").SendD(c)
		return
	}

	var user models.User

	if strings.ToLower(request.AuthType) == AUTH_GOOGLE_APP {

		err = db.Model(&models.User{}).Where("google = ?", request.Email).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

				// User does not exist, so create a new user
				user = models.User{
					Google:       &request.Email,
					UserId:       uuid.NewString(),
					ReferralCode: referral.GetReferalCode(),
				}
				err = db.Model(&models.User{}).Create(&user).Error
				if err != nil {
					logwrapper.Errorf("failed to create user: %s", err)
					httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error : "+err.Error()).SendD(c)
					return
				}
				meta := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
				go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Created, Metadata: &meta})

				// if user.ReferralCode == "" {
				// 	referral.GenerateReferralCodeForUser(user)
				// }
			} else {
				// Other error occurred
				logwrapper.Errorf("failed to retrieve user: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error : "+err.Error()).SendD(c)
				return
			}
		}

	} else if strings.ToLower(request.AuthType) == AUTH_APPLE_APP {
		err = db.Model(&models.User{}).Where("apple_id = ?", request.AppleID).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// User does not exist, so create a new user

				var email *string
				if request.Email == "" {
					email = nil // This will store NULL in the database
				} else {
					email = &request.Email // Store the provided email
				}

				user = models.User{
					Apple:        email,
					UserId:       uuid.NewString(),
					AppleId:      &request.AppleID,
					ReferralCode: referral.GetReferalCode(),
				}
				err = db.Model(&models.User{}).Create(&user).Error
				if err != nil {
					logwrapper.Errorf("failed to create user: %s", err)
					httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error : "+err.Error()).SendD(c)
					return
				} else {
					meta := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
					go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Created, Metadata: &meta})
					logwrapper.Infof("user created successfully")
				}
			} else {
				// Other error occurred
				logwrapper.Errorf("failed to retrieve user: %s", err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error : "+err.Error()).SendD(c)
				return

			}
		}

	}
	c.Set(paseto.CTX_USER_ID, user.UserId)

	if user.ReferralCode == "" {
		referral.GenerateReferralCodeForUser(user)
	}

	customClaims := claims.NewWithEmail(user.UserId, user.Email)
	pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
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
	db := database.GetDb()
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
			meta := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
			go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Created, Metadata: &meta})

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
				meta := "email  : " + *user.Email + ", wallet address : " + *user.WalletAddress
				go useractivity.Save(models.UserActivity{UserId: user.UserId, Modules: module.Account, Action: actions.Updated, Metadata: &meta})

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
	// pvKey, err := hex.DecodeString(load.Cfg.PASETO_PRIVATE_KEY[2:])
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

	if user.ReferralCode == "" {
		referral.GenerateReferralCodeForUser(user)
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
	db := database.GetDb() // Initialize the database connection
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
