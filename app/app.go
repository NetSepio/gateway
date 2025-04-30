package app

import (
	"strings"
	"time"

	"github.com/NetSepio/gateway/api"
	"github.com/NetSepio/gateway/app/routines/reportroutine"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/stripe/stripe-go/v76"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/config/redisconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	redisconfig.InitRedis()
	dbconfig.Migrate()
	stripe.Key = envconfig.EnvVars.STRIPE_SECRET_KEY
	logwrapper.Init()

	GinApp = gin.Default()

	if strings.ToLower(envconfig.EnvVars.GIN_MODE) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin.SetMode(gin.ReleaseMode)

	corsM := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"*"},
	})
	GinApp.Use(corsM)

	api.ApplyRoutes(GinApp)

	//adding health check

	GinApp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	go reportroutine.StartProcessingReportsPeriodically()
	// go webreview.Init()
}
