package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

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
	Config.CertificatePath = "/home/magesh/tmp/ca/"
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
func GetNewNxoService() NxoServiceIntf {
	return &NxoService{
		FacadeMap: make(map[string]string),
	}
}

// ------------------------------------------------------------------------
// Get Facade Endpoints
// ------------------------------------------------------------------------
func (h *NxoService) CallFacade(r *http.Request) error {
	// Construct the CallFacade API URL
	facadeURL := h.FacadeURL

	// Modify request headers and body if needed for CallFacade
	req, err := http.NewRequest(http.MethodGet, facadeURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create CallFacade request: %v", err)
	}

	//set the credentials and headers for calling facade api
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+h.FacadeCred)

	// HTTPS client <FIXME> try to reuse client
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: h.NxoConfig.TlsClientConfig,
		},
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("CallFacade API failed: %v", err)
	}

	// Read and return the response body (as-is)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read CallFacade response body: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("CallFacade API returned status: %s", resp.Status)
	}
	//Parse and update the FacadeMap
	if err := json.Unmarshal(body, &h.FacadeMap); err != nil {
		return fmt.Errorf("failed to parse CallFacade response JSON: %v", err)
	}
	log.Printf("Facade Endpoints: %v", h.FacadeMap)
	return nil
}

// ------------------------------------------------------------------------
// Call iDRAC Endpoints
// ------------------------------------------------------------------------
func (h *NxoService) CalliDRAC(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	proxyURL := h.EOProxyURL + r.URL.Path
	maxRetry := 3

	// --- Prepare body (for POST/PUT etc). Read once and reuse ---
	var bodyBytes []byte
	if r.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %v", err)
		}
		r.Body.Close()
	}

	// HTTPS client <FIXME> try to reuse client
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: h.NxoConfig.TlsClientConfig,
		},
	}

	// Retry logic
	for attempt := 1; attempt <= maxRetry; attempt++ {
		// Recreate body reader each retry (safe for multiple sends)
		var reqBody io.Reader
		if bodyBytes != nil {
			reqBody = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequest(r.Method, proxyURL, reqBody)
		if err != nil {
			return nil, fmt.Errorf("failed to create iDRAC proxy request: %v", err)
		}

		// Copy original headers
		for key, values := range r.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
		// Set iDRAC-specific headers
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Basic "+h.EOProxyCred)

		// Execute request
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Attempt %d: request failed: %v", attempt, err)
			time.Sleep(time.Duration(attempt*attempt) * time.Second)
			continue
		}

		// Always close body
		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr != nil {
			return nil, fmt.Errorf("failed to read iDRAC response: %v", readErr)
		}

		// Retry on non-200
		if resp.StatusCode != http.StatusOK {
			log.Printf("Attempt %d: iDRAC returned %s - body: %s", attempt, resp.Status, string(body))
			// Refresh FacadeMap
			if err := h.CallFacade(r); err != nil {
				log.Printf("Error refreshing FacadeMap: %v", err)
			}
			time.Sleep(time.Duration(attempt*attempt) * time.Second)
			continue
		}
		// Success
		return body, nil
	}

	return nil, fmt.Errorf("iDRAC request failed after %d attempts", maxRetry)
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

	//on prem, both facade and proxy are within DSPO
	h.FacadeCred = getenv("NXO_FACADE_CRED", "admin")
	h.FacadeURL = getenv("NXO_FACADE_URL", "https://localhost:9000")
	h.EOProxyCred = getenv("NXO_EOPROXY_CRED", "admin")
	h.EOProxyURL = getenv("NXO_EOPROXY_URL", "https://localhost:9000")

	//SaaS, facade url is within DSPO , iDRAC proxy is on a different url.
	if strings.ToLower(installType) != "onprem" {
		h.EOProxyCred = getenv("NXO_EOPROXY_CRED", "admin")
		h.EOProxyURL = getenv("NXO_EOPROXY_URL", "https://localhost:9001")
	}

	//Plugin Service Port
	h.PLUGINS_PORT = getenv("NXO_PLUGINS_PORT", "8443")

	//Config
	h.NxoConfig = GetConfig()
	return nil
}

// ------------------------------------------------------------------------
// Start the service
// ------------------------------------------------------------------------
func (h *NxoService) Start() error {
	addr := fmt.Sprintf("%s:%s", "0.0.0.0", h.PLUGINS_PORT)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello from service on %s!\n", addr)
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		type result struct {
			data []byte
			err  error
		}

		resultChan := make(chan result, 1)

		go func() {
			resp, err := h.CalliDRAC(w, r.WithContext(ctx))
			resultChan <- result{data: resp, err: err}
		}()

		select {
		case <-ctx.Done():
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			return
		case res := <-resultChan:
			if res.err != nil {
				http.Error(w, fmt.Sprintf("Error: %v", res.err), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(res.data)
			return
		}
	})

	h.Server = &http.Server{
		Addr:              addr,
		Handler:           mux,
		TLSConfig:         h.NxoConfig.TLSServerConfig,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("Service starting on %s", addr)
	return h.Server.ListenAndServeTLS(h.NxoConfig.ServerCert, h.NxoConfig.ServerKey)
}

// ------------------------------------------------------------------------
// Start the service
// ------------------------------------------------------------------------
func (h *NxoService) Stop() error {
	//Stop the service
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := h.Server.Shutdown(ctx); err != nil {
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
