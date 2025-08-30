package models

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	utils "golang-examples/nxo-dap-plugins/plugins/nxo-plugin/utils"
)

var (
	Config *NxoConfig
	mu     sync.Mutex
)

// ------------------------------------------------------------------------
// Get Configuration
// ------------------------------------------------------------------------
func GetConfig() *NxoConfig {
	mu.Lock()
	defer mu.Unlock()
	if Config == nil {
		Config = &NxoConfig{}
	}
	//<FIXME> - Change the path to env variable or config map
	Config.CertificatePath = "/tmp/ca/"
	Config.CaCert = Config.CertificatePath + "ca.crt"
	Config.ClientCert = Config.CertificatePath + "client.crt"
	Config.ClientKey = Config.CertificatePath + "client.key"
	Config.ServerCert = Config.CertificatePath + "server.crt"
	Config.ServerKey = Config.CertificatePath + "server.key"

	Config.TLSVerify = false
	if val, ok := os.LookupEnv("NXO_TLS_VERIFY"); ok && (strings.ToLower(val) == "true") {
		Config.TLSVerify = true
	}
	serverConfig, err := utils.GetTLServerConfig(Config.ServerCert, Config.ServerKey, Config.CaCert, Config.TLSVerify)
	if err != nil {
		log.Fatalf("Error getting TLS server config: %v", err)
	}

	Config.TLSServerConfig = serverConfig
	clientConfig, err := utils.GetTlsClientConfig(Config.ClientCert, Config.ClientKey, Config.CaCert, Config.TLSVerify)
	if err != nil {
		log.Fatalf("Error getting TLS client config: %v", err)
	}
	Config.TlsClientConfig = clientConfig
	//Config.TLSConfig = nil
	return Config
}

// ------------------------------------------------------------------------
// Create a new service instance
// ------------------------------------------------------------------------
func GetNewNxoService(installType string, cred string, host string, port string) *NxoService {
	return &NxoService{
		InstallType: installType,
		OrgCred:     cred,
		OrgHost:     host,
		OrgPort:     port,
	}
}

// ------------------------------------------------------------------------
// Get Facade Endpoints
// ------------------------------------------------------------------------
func (h *NxoService) CallFacade(ctx context.Context) error {
	//Get Endpoints
	return nil
}

// ------------------------------------------------------------------------
// Call iDRAC Endpoints
// ------------------------------------------------------------------------
func (h *NxoService) CalliDRAC(ctx context.Context) error {
	//Call iDRAC
	return nil
}

// ------------------------------------------------------------------------
// Initialize the service
// ------------------------------------------------------------------------
func (h *NxoService) Init() (error error) {
	defer func() {
		if r := recover(); r != nil {
			error = fmt.Errorf("error during initialization process: %v", r)
		}
	}()

	//Initialize
	installType := getenv("NXO_INSTALL_TYPE", "onprem")
	h.InstallType = installType

	//on prem
	h.OrgCred = getenv("NXO_ORG_CRED", "admin")
	//h.OrgHost = getenv("NXO_ORG_HOST", "nxo-organization")
	//h.OrgPort = getenv("NXO_ORG_PORT", "443")
	h.FacadeURL = getenv("NXO_FACADE_URL", fmt.Sprintf("http://%s:%s", h.OrgHost, h.OrgPort))
	h.ProxyCred = getenv("NXO_ORG_CRED", "admin")

	if strings.ToLower(installType) != "onprem" {
		h.ProxyCred = getenv("NXO_PROXY_CRED", "admin")
	}
	h.NxoConfig = GetConfig()
	return nil
}

// ------------------------------------------------------------------------
// Start the service
// ------------------------------------------------------------------------
func (h *NxoService) Start() error {
	addr := fmt.Sprintf("%s:%s", h.OrgHost, h.OrgPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service on %s!", addr)
	})

	h.Server = &http.Server{
		Addr:      addr,
		Handler:   mux,
		TLSConfig: h.NxoConfig.TLSServerConfig,
	}

	log.Printf("Service starting on %s", addr)
	return h.Server.ListenAndServeTLS(h.NxoConfig.ServerCert, h.NxoConfig.ServerKey)
}

// ------------------------------------------------------------------------
// Start the service
// ------------------------------------------------------------------------
func (h *NxoService) Stop() error {
	//Stop the service
	if err := h.Server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error shutting down server: %v", err)
	}
	return nil
}

// ------------------------------------------------------------------------
// Get env variable or default value
// ------------------------------------------------------------------------
func getenv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
