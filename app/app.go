package app

import (
	"time"

	"github.com/NetSepio/gateway/api"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/config/netsepio"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	logwrapper.Init()
	netsepio.InitRolesId()

	GinApp = gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
		AllowOrigins:     envconfig.EnvVars.ALLOWED_ORIGIN})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
	// go webreview.Init()
}
