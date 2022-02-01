package app

import (
	"github.com/TheLazarusNetwork/marketplace-engine/api"
	"github.com/TheLazarusNetwork/marketplace-engine/config"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/marketplace-engine/config/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init(envPath string) {
	config.Init(envPath)
	logwrapper.Init()
	creatify.InitRolesId()
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
}
