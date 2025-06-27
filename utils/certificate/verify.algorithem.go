package certificate

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"go.uber.org/zap"
	"github.com/NetSepio/gateway/utils/load"
)

func verifyCertFromFiles(caCertBase64, caKeyBase64 string, certPEM []byte) error {

	// Decode CA
	caCertPEM, err := base64.StdEncoding.DecodeString(caCertBase64)
	if err != nil {
		// Use logger instead of gin context
		load.Logger.Error("Invalid CA cert in env", zap.Error(err))
		return errors.New("invalid CA cert in env")
	}
	caKeyPEM, err := base64.StdEncoding.DecodeString(caKeyBase64)
	if err != nil {

		load.Logger.Error("Invalid CA key in env", zap.Error(err))
		return errors.New("invalid CA key in env")
	}

	// Parse CA
	caCert, _, err := parseCACertKey(caCertPEM, caKeyPEM)
	if err != nil {
		load.Logger.Error("Failed to parse CA cert/key", zap.Error(err))
		return errors.New("failed to parse CA cert/key")
	}

	// Decode and parse uploaded cert
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		load.Logger.Error("Invalid certificate PEM format")
		return errors.New("invalid certificate PEM format")
	}

	leafCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		load.Logger.Error("Failed to parse certificate", zap.Error(err))
		return errors.New("failed to parse certificate")
	}

	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	if _, err := leafCert.Verify(x509.VerifyOptions{Roots: roots}); err != nil {
		load.Logger.Error("Certificate verification failed", zap.Error(err))
		return errors.New("certificate verification failed: " + err.Error())
	}
	return nil
}
func parseCACertKey(certPEM, keyPEM []byte) (*x509.Certificate, ed25519.PrivateKey, error) {
	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		return nil, nil, errors.New("invalid CA certificate PEM")
	}
	caCert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil || keyBlock.Type != "PRIVATE KEY" {
		return nil, nil, errors.New("invalid CA private key PEM")
	}
	keyParsed, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	edPrivKey, ok := keyParsed.(ed25519.PrivateKey)
	if !ok {
		return nil, nil, errors.New("CA private key is not Ed25519")
	}

	return caCert, edPrivKey, nil
}
