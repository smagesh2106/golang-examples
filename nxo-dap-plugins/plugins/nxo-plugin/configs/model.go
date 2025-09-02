package configs

import (
	"crypto/tls"
	"sync"
)

var (
	config *NxoConfig
	mu     sync.Once
)

type NxoConfig struct {
	CaCert     string
	ClientCert string
	ClientKey  string
	ServerCert string
	ServerKey  string
	TLSVerify  bool
	//TLSConfig 		 func() (*tls.Config, error)
	TLSServerConfig *tls.Config
	TlsClientConfig *tls.Config

	//Server properties
	ReadHeaderTimeout int
	ReadTimeout       int
	WriteTimeout      int
	IdleTimeout       int
	HttpClientTimeout int

	//Facade / EOProxy / SaaS Properties
	FacadeURL      string
	FacadeCred     string
	FacadeAuthType string

	EOProxyURL      string
	EOProxyCred     string
	EOProxyAuthType string

	SaaSProxyURL      string
	SaaSProxyCred     string
	SaaSProxyAuthType string

	//Server port
	PLUGINS_PORT string
	//Install Type onprem / SaaS
	InstallType string
}
