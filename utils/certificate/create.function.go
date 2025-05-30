package certificate

import (
	"encoding/base64"
	"fmt"

	"go.uber.org/zap"
	"netsepio-gateway-v1.1/utils/load"
)

func GenerateCertificateAndReturn(domain string) ([]byte, []byte, error) {

	caCertBase64 := load.Cfg.ROOT_CA_CERT_BASE64
	caKeyBase64 := load.Cfg.ROOT_CA_KEY_BASE64

	fmt.Println("CA Cert Base64:", caCertBase64)
	fmt.Println("CA Key Base64:", caKeyBase64)

	// Decode base64-encoded CA cert and key
	caCertPEM, err := base64.StdEncoding.DecodeString(caCertBase64)
	if err != nil {
		// Return a formatted error message if decoding fails
		// Use logger instead of gin context
		load.Logger.Error("Failed to decode ROOT_CA_CERT_BASE64", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to decode ROOT_CA_CERT_BASE64: %v", err)
	}
	caKeyPEM, err := base64.StdEncoding.DecodeString(caKeyBase64)
	if err != nil {
		// Return a formatted error message if decoding fails
		load.Logger.Error("Failed to decode ROOT_CA_KEY_BASE64", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to decode ROOT_CA_KEY_BASE64: %v", err)
	}

	// Parse CA cert and private key
	caCert, caKey, err := parseCACertKey(caCertPEM, caKeyPEM)
	if err != nil {
		// Return a formatted error message if parsing fails
		load.Logger.Error("Failed to parse CA cert/key", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to parse CA cert/key: %v", err)
	}

	// Generate new leaf certificate and key
	leafCertPEM, leafKeyPEM, err := generateSignedCert(caCert, caKey, domain)
	if err != nil {
		// Return a formatted error message if certificate generation fails
		load.Logger.Error("Failed to generate leaf certificate", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to generate leaf certificate: %v", err)
	}

	return leafCertPEM, leafKeyPEM, nil
}
