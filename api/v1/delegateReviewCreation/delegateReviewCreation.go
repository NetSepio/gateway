package delegatereviewcreation

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
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
		httphelper.ErrResponse(c, http.StatusBadRequest, "payload is invalid")
		return
	}
	//TODO function id and gas from env
	command := fmt.Sprintf("move run --function-id %s::netsepio::delegate_submit_review --max-gas %d --gas-unit-price %d --args", "0x1da41025906f10f17f74f6c6851cb3d192acdd31131123f67e800aa5358b5bc1", 3046, 100)
	args := append(strings.Split(command, " "),
		argA(request.Voter), argS(request.MetaDataUri), argS(request.Category), argS(request.DomainAddress), argS(request.SiteUrl), argS(request.SiteType), argS(request.SiteTag), argS(request.SiteSafety), argS(request.SiteIpfsHash))
	cmd := exec.Command("./aptos", args...)
	fmt.Println(strings.Join(args, " "))
	// The `Output` method executes the command and
	// collects the output, returning its value
	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			httphelper.NewInternalServerError(c, "failed to call %v of %v, error: %v %s %s", "delegate_submit_review", "NETSEPIO", err.Error(), err.Stderr, o)
			return
		}
		httphelper.NewInternalServerError(c, "failed to call %v of %v, error: %v %s", "delegate_submit_review", "NETSEPIO", err.Error(), o)
		return
	}

	txResult, err := UnmarshalTxResult(o)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get transaction result")
		return
	}
	payload := DelegateReviewCreationPayload{
		TransactionVersion: txResult.Result.Version,
		TransactionHash:    txResult.Result.TransactionHash,
	}
	logwrapper.Infof("tx is %v", txResult)
	httphelper.SuccessResponse(c, "request successfully send, review will be delegated soon", payload)
}
