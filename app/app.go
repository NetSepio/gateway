package app

import (
	"time"

	"github.com/NetSepio/gateway/api"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/stripe/stripe-go/v76"

	"github.com/NetSepio/gateway/config/constants"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	constants.InitConstants()
	logwrapper.Init()
	stripe.Key = envconfig.EnvVars.STRIPE_SECRET_KEY

	GinApp = gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     envconfig.EnvVars.ALLOWED_ORIGIN})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
	// go webreview.Init()
}
