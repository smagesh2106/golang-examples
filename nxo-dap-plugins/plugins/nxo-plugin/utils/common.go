package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// ------------------------------------------------------------------------
// Basic TLS Config
// ------------------------------------------------------------------------
func basicConfig(certPath string, keyPath string, caCertPath string) (*tls.Certificate, *x509.CertPool, error) {

	// Validate file paths
	if _, err := os.Stat(certPath); err != nil {
		return nil, nil, fmt.Errorf("client certificate file not found: %v", err)
	}
	if _, err := os.Stat(keyPath); err != nil {
		return nil, nil, fmt.Errorf("client key file not found: %v", err)
	}
	if _, err := os.Stat(caCertPath); err != nil {
		return nil, nil, fmt.Errorf("CA certificate file not found: %v", err)
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to load key pair: %v", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to load CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		return nil, nil, fmt.Errorf("Failed to append CA cert: %v", err)
	}
	return &cert, caCertPool, nil
}

// ------------------------------------------------------------------------
// Get TLS Server Config
// ------------------------------------------------------------------------
func GetTLServerConfig(cert string, key string, cacert string, verify bool) (*tls.Config, error) {
	fmt.Println("Creating TLS server configuration...")
	tlsCert, caCertPool, err := basicConfig(cert, key, cacert)
	fmt.Println("Basic Server TLS configuration created.")
	if err != nil {
		return nil, fmt.Errorf("error in ServerConfig, server certs,key,ca missing: %v", err)
	}

	verifyFlag := tls.NoClientCert
	if verify {
		verifyFlag = tls.RequireAndVerifyClientCert
	}
	fmt.Println("Client certificate verification set.")
	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		MaxVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{*tlsCert},
		ClientAuth:   verifyFlag,
		ClientCAs:    caCertPool,
	}
	//tlsConfig.BuildNameToCertificate()
	fmt.Println("TLS server configuration created.")
	return tlsConfig, nil
}

// ------------------------------------------------------------------------
// Get TLS Client Config
// ------------------------------------------------------------------------
func GetTlsClientConfig(cert string, key string, cacert string, _ bool) (*tls.Config, error) {
	tlsCert, caCertPool, err := basicConfig(cert, key, cacert)
	if err != nil {
		return nil, fmt.Errorf("error in ServerConfig, client certs,key,ca missing: %v", err)
	}

	// Configure TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*tlsCert},
		RootCAs:      caCertPool,
	}
	return tlsConfig, nil
}

func GetENV(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
