package app

import (
	"netsepio-api/api"
	"netsepio-api/db"
	"netsepio-api/util/pkg/logwrapper"

	"netsepio-api/config"

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
