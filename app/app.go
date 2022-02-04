package app

import (
	"github.com/TheLazarusNetwork/netsepio-engine/api"
	"github.com/TheLazarusNetwork/netsepio-engine/app/routines/webreview"
	"github.com/TheLazarusNetwork/netsepio-engine/config"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/config/netsepio"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init(envPath string, logBasePath string) {
	config.Init(envPath)
	logwrapper.Init(logBasePath)
	netsepio.InitRolesId()
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
	go webreview.Init()
}
