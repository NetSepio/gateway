package certificate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/NetSepio/gateway/utils/certificate"
	"github.com/NetSepio/gateway/utils/load"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/proxy")
	{
		g.GET("/auth/:domain", generateCertificate)
		g.POST("/verify", VerifyCertificate)
	}

}

func generateCertificate(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		load.Logger.Error("Missing domain parameter in query")
		c.JSON(400, gin.H{"error": "Missing domain parameter"})
		return
	}

	certPEM, _, err := certificate.GenerateCertificateAndReturn(domain)
	if err != nil {
		load.Logger.Error("Certificate generation failed", zap.Error(err))
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Send certificate.pem as downloadable file
	c.Header("Content-Disposition", "attachment; filename=certificate.pem")
	c.Data(http.StatusOK, "application/x-pem-file", certPEM)
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
