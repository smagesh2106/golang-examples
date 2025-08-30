package models

import (
	"context"
	"crypto/tls"
	"net/http"
)

type NxoServiceIntf interface {
	Init() error
	Start() error
	Stop()

	CallFacade(ctx context.Context) error
	CalliDRAC(ctx context.Context) error
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
	OrgCred   string `json:"orgcred" yaml:"orgcred"`
	OrgHost   string `json:"orghost" yaml:"orghost"`
	OrgPort   string `json:"orgport" yaml:"orgport"`
	FacadeURL string `json:"facadeurl" yaml:"facadeurl"`

	//SaaS Model
	ProxyCred string `json:"proxycred" yaml:"proxycred"`

	//Config
	NxoConfig *NxoConfig `json:"nxoconfig" yaml:"nxoconfig"`

	Server *http.Server
}
