package certificate

import (
	"mime/multipart"

	"go.uber.org/zap"
	"netsepio-gateway-v1.1/utils/load"
)

func VerifyCertificate(file *multipart.FileHeader) (bool, error) {

	src, err := file.Open()
	if err != nil {
		load.Logger.Error("Failed to open uploaded file", zap.Error(err))
		return false, err
	}
	defer src.Close()

	// Read file into memory
	certPEM := make([]byte, file.Size)
	_, err = src.Read(certPEM)
	if err != nil {
		load.Logger.Error("Failed to read uploaded file", zap.Error(err))
		return false, err
	}

	caCertBase64 := load.Cfg.ROOT_CA_CERT_BASE64
	caKeyBase64 := load.Cfg.ROOT_CA_KEY_BASE64

	// Verify leaf cert
	if err := verifyCertFromFiles(caCertBase64, caKeyBase64, certPEM); err != nil {
		load.Logger.Error("Verification failed", zap.Error(err))
		return false, err
	}

	load.Logger.Info("Leaf certificate verified successfully!")
	return true, nil

}
