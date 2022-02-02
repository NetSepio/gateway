package authenticate

import (
	"net/http"

	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/models"
	"github.com/TheLazarusNetwork/netsepio-engine/models/claims"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/auth"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/cryptosign"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"

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

	db := dbconfig.GetDb()
	//TODO remove flow id if 200
	var req AuthenticateRequest
	c.BindJSON(&req)

	//Get flowid type
	var flowIdData models.FlowId
	err := db.Model(&models.FlowId{}).Where("flow_id = ?", req.FlowId).First(&flowIdData).Error
	if err != nil {
		logwrapper.Error(err)
		httphelper.ErrResponse(c, 500, "Unexpected error occured")
		return
	}

	if flowIdData.FlowIdType != models.AUTH {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Flow id not created for auth")
		return
	}

	if err != nil {
		logwrapper.Error(err)
		httphelper.ErrResponse(c, 500, "Unexpected error occured")
		return
	}
	userAuthEULA := envutil.MustGetEnv("AUTH_EULA")
	message := userAuthEULA + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, "Flow Id not found")
		return
	}

	if err != nil {
		logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		jwtPrivateKey := envutil.MustGetEnv("JWT_PRIVATE_KEY")
		jwtToken, err := auth.GenerateToken(customClaims, jwtPrivateKey)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to generate token, error %v", err.Error())
			return
		}
		db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{})
		payload := AuthenticatePayload{
			Token: jwtToken,
		}
		httphelper.SuccessResponse(c, "Token generated successfully", payload)
	} else {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}
