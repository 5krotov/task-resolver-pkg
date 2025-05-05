package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"os"
)

func loadTLSCerts(certPath, keyPath, caPath string) (*x509.CertPool, *tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load server cert: %v", err)
	}

	caPEM, err := os.ReadFile(caPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CA cert: %v", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPEM) {
		return nil, nil, fmt.Errorf("failed to add CA cert")
	}

	return certPool, &cert, nil
}

func loadTLSServerCreds(certPath, keyPath, caPath string) (credentials.TransportCredentials, error) {
	certPool, cert, err := loadTLSCerts(certPath, keyPath, caPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load certs: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
	}

	return credentials.NewTLS(config), nil
}

func loadTLSClientCreds(certPath, keyPath, caPath, serverName string) (credentials.TransportCredentials, error) {
	certPool, cert, err := loadTLSCerts(certPath, keyPath, caPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load certs: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		RootCAs:      certPool,
		ServerName:   serverName,
		MinVersion:   tls.VersionTLS13,
	}

	return credentials.NewTLS(config), nil
}
