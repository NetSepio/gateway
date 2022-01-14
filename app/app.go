package app

import (
	"github.com/TheLazarusNetwork/marketplace-engine/api"
	"github.com/TheLazarusNetwork/marketplace-engine/db"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/marketplace-engine/config"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	config.Init()
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	db.InitDB()
	logwrapper.Init()
}
