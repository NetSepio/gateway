package delegatereviewcreation

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateReviewCreation")
	{
		g.Use(paseto.PASETO)
		g.POST("", deletegateReviewCreation)
	}
}

func argS(s string) string {
	return "string:" + s
}

func argA(s string) string {
	return "address:" + s
}

func deletegateReviewCreation(c *gin.Context) {
	var request DelegateReviewCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, "payload is invalid").SendD(c)
		return
	}
	command := fmt.Sprintf("move run --function-id %s::netsepio::delegate_submit_review --max-gas %d --gas-unit-price %d --args", envconfig.EnvVars.FUNCTION_ID, envconfig.EnvVars.GAS_UNITS, envconfig.EnvVars.GAS_PRICE)
	args := append(strings.Split(command, " "),
		argA(request.Voter), argS(request.MetaDataUri), argS(request.Category), argS(request.DomainAddress), argS(request.SiteUrl), argS(request.SiteType), argS(request.SiteTag), argS(request.SiteSafety), argS(request.SiteIpfsHash))
	cmd := exec.Command("./aptos", args...)
	fmt.Println(strings.Join(args, " "))
	// The `Output` method executes the command and
	// collects the output, returning its value
	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			logwrapper.Errorf("failed to call %v of %v, error: %v %s %s", "delegate_submit_review", "NETSEPIO", err.Error(), err.Stderr, o)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to call %v of %v, error: %v %s", "delegate_submit_review", "NETSEPIO", err.Error(), o)
		return
	}

	txResult, err := UnmarshalTxResult(o)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		logwrapper.Errorf("failed to get transaction result")
		return
	}
	payload := DelegateReviewCreationPayload{
		TransactionVersion: txResult.Result.Version,
		TransactionHash:    txResult.Result.TransactionHash,
	}
	logwrapper.Infof("tx is %v", txResult)
	httpo.NewSuccessResponseP(200, "request successfully send, review will be delegated soon", payload).SendD(c)
}
