package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NetSepio/gateway/internal/api"
	"github.com/NetSepio/gateway/utils/load"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {

	if strings.ToLower(load.Cfg.GIN_MODE) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	corsM := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     load.Cfg.ALLOWED_ORIGIN,
	})

	GinApp = gin.Default()

	GinApp.Use(corsM)
	GinApp.Use(gin.Recovery())
	GinApp.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		c.Abort()
	}))
	GinApp.Use(gin.ErrorLogger())
	// Removed gin.ErrorLoggerWithConfig as it does not exist in the Gin framework
	GinApp.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] %s \"%s %s %s %d %s\" %s\n",
				param.TimeStamp.Format(time.RFC3339),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.ErrorMessage,
			)
		},
		Output:    nil,
		SkipPaths: []string{"/health", "/metrics"},
	}))

	GinApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Api Not Found",
		})
	})
	GinApp.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"error": "Method Not Allowed",
		})
	})

	GinApp.Use(gin.Recovery())

	api.ApplyRoutes(GinApp)

}

func Start() {
	if err := GinApp.Run(":" + strconv.Itoa(load.Cfg.APP_PORT)); err != nil {
		panic(err)
	}
}
