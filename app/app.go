package app

import (
	"time"

	"github.com/TheLazarusNetwork/netsepio-engine/api"
	"github.com/TheLazarusNetwork/netsepio-engine/config"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/config/netsepio"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init(envPath string, logBasePath string) {
	config.Init(envPath)
	logwrapper.Init(logBasePath)
	netsepio.InitRolesId()

	GinApp = gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD","OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type","Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{envutil.MustGetEnv("ALLOWED_ORIGIN")}})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
}
