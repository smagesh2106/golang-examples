package models

import (
	"crypto/tls"
	"net/http"
)

type NxoServiceIntf interface {
	Init() error
	Start() error
	Stop() error

	CallFacade(r *http.Request) error
	CalliDRAC(w http.ResponseWriter, r *http.Request) ([]byte, error)
}

type NxoConfig struct {
	CertificatePath string
	CaCert          string
	ClientCert      string
	ClientKey       string
	ServerCert      string
	ServerKey       string
	TLSVerify       bool
	//TLSConfig 		 func() (*tls.Config, error)
	TLSServerConfig *tls.Config
	TlsClientConfig *tls.Config
}

type NxoService struct {
	// install type onprem / SaaS
	InstallType string `json:"installtype" yaml:"installtype"`

	//On Prem Model
	FacadeCred string `json:"orgcred" yaml:"orgcred"`
	FacadeURL  string `json:"facadeurl" yaml:"facadeurl"`

	//SaaS Model
	EOProxyURL  string `json:"eoproxyurl" yaml:"eoproxyurl"`
	EOProxyCred string `json:"proxycred" yaml:"proxycred"`

	//Plugin Service Port
	PLUGINS_PORT string `json:"pluginsport" yaml:"pluginsport"`
	//Config
	NxoConfig *NxoConfig `json:"nxoconfig" yaml:"nxoconfig"`

	// Facade Endpoints map
	FacadeMap map[string]string `json:"facademap" yaml:"facademap"`

	//HTTP Server
	Server *http.Server
}
