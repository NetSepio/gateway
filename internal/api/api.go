package api

import (
	"github.com/gin-gonic/gin"
	v11 "github.com/NetSepio/gateway/internal/api/v.1.1"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v11.ApplyRoutes(api)
		v11.ApplyRoutesV1_1(api)
	}
}
