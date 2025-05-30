package certificate

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"netsepio-gateway-v1.1/utils/certificate"
	"netsepio-gateway-v1.1/utils/load"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/certificate")
	{
		g.POST("/verify", VerifyCertificate)
	}

}

func VerifyCertificate(c *gin.Context) {
	file, err := c.FormFile("cert")
	if err != nil {
		// add error logging from utils/load/load.logger.go
		load.Logger.Error("Failed to get 'cert' field from form", zap.Error(err))
		c.JSON(400, gin.H{"error": "Missing certificate file in 'cert' field"})
		return
	}
	valid, err := certificate.VerifyCertificate(file)
	if err != nil {
		load.Logger.Error("Certificate verification failed", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid certificate"})
		return
	}
	if !valid {
		load.Logger.Warn("Certificate is not valid", zap.String("filename", file.Filename))
		c.JSON(400, gin.H{"error": "Certificate is not valid"})
		return
	} else {
		c.JSON(200, gin.H{"message": "Certificate is valid"})
		return
	}

}
