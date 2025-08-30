package utils

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

func basicConfig(certPath string, keyPath string, caCertPath string) (*tls.Certificate, *x509.CertPool) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal("Failed to load server key pair:", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatal("Failed to read CA cert:", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA cert")
	}
	return &cert, caCertPool
}

func GetTLServerConfig(path string, verify bool) (*tls.Config, error) {
	//FIXME validate if the path and files exist, use the verify flag
	cert, caCertPool := basicConfig(path+"/server.crt", path+"/server.key", path+"/ca.crt")

	verifyFlag := tls.NoClientCert
	if verify {
		verifyFlag = tls.RequireAndVerifyClientCert
	}
	// Configure TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientAuth:   verifyFlag,
		ClientCAs:    caCertPool,
	}
	//tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}

func GetTlsClientConfig(path string, _ bool) (*tls.Config, error) {
	//FIXME validate if the path and files exist, use the verify flag
	cert, caCertPool := basicConfig(path+"/client.crt", path+"/client.key", path+"/ca.crt")

	// Configure TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		RootCAs:      caCertPool,
	}
	return tlsConfig, nil
}
