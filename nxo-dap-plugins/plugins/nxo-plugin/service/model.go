package service

import (
	"golang-examples/nxo-dap-plugins/plugins/nxo-plugin/configs"
	"net/http"
)

type NxoServiceIntf interface {
	Init() error
	Start() error
	Stop() error

	CallFacade(r *http.Request) error
	CalliDRAC(w http.ResponseWriter, r *http.Request) ([]byte, error)
}

type NxoService struct {
	//Config
	NxoConfig *configs.NxoConfig `json:"nxoconfig" yaml:"nxoconfig"`

	// Facade Endpoints map
	FacadeMap map[string]string `json:"facademap" yaml:"facademap"`

	//HTTP Server
	Server *http.Server
}
