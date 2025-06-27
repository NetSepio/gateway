package ctx

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/utils/status"
)

func GetCreatorID(c *gin.Context) (string, string, error) {
	if id := c.GetString(paseto.CTX_USER_ID); id != "" {
		return id, paseto.CTX_USER_ID, nil
	}
	if id := c.GetString(paseto.CTX_ORGANISATION_ID); id != "" {
		return id, paseto.CTX_ORGANISATION_ID, nil
	}
	if id := c.GetString(paseto.CTX_ORG_APP_ID); id != "" {
		return id, paseto.CTX_ORG_APP_ID, nil
	}
	return "", status.INVALID, errors.New("user or organisation id not found in context")
}
