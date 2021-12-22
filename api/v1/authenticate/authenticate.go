package authenticate

import (
	"net/http"
	"netsepio-api/db"
	"netsepio-api/models"
	"netsepio-api/models/claims"
	"netsepio-api/util/pkg/auth"
	"netsepio-api/util/pkg/cryptosign"
	"netsepio-api/util/pkg/httphelper"
	"netsepio-api/util/pkg/logwrapper"
	"os"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {

	//TODO remove flow id if 200
	var req AuthenticateRequest
	c.BindJSON(&req)

	var role models.Role
	var defaultRoleId = 1
	err := db.Db.Model(&models.Role{}).First(&role, defaultRoleId).Error
	if err != nil {
		logwrapper.Log.Error(err)
		httphelper.ErrResponse(c, 500, "Unexpected error occured")
		return
	}

	message := role.Eula + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, "Flow Id not found")
		return
	}

	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
		jwtToken, err := auth.GenerateToken(customClaims, jwtPrivateKey)
		if err != nil {
			httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
			return
		}
		payload := AuthenticatePayload{
			Token: jwtToken,
		}
		httphelper.SuccessResponse(c, "Token generated successfully", payload)
	} else {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}
