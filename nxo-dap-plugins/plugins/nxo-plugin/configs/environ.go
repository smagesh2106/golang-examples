package configs

import (
	"fmt"
	"golang-examples/nxo-dap-plugins/plugins/nxo-plugin/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------
// Get Configuration
// ------------------------------------------------------------------------
func GetConfig() *NxoConfig {
	fmt.Println("Loading configuration from environment variables...")
	mu.Do(func() {
		fmt.Println("Initializing configuration...")
		c := &NxoConfig{}
		//Certificate related paths
		c.CaCert = utils.GetENV("CA_CERT", "/home/magesh/tmp/ca/ca.crt")
		c.ClientCert = utils.GetENV("NXO_CLIENT_CERT", "/home/magesh/tmp/ca/client.crt")
		c.ClientKey = utils.GetENV("NXO_CLIENT_KEY", "/home/magesh/tmp/ca/client.key")
		c.ServerCert = utils.GetENV("NXO_CLIENT_CERT", "/home/magesh/tmp/ca/server.crt")
		c.ServerKey = utils.GetENV("NXO_CLIENT_KEY", "/home/magesh/tmp/ca/server.key")
		fmt.Println("Certificate paths set.")
		//TLS Verify
		c.TLSVerify = false
		if val, ok := os.LookupEnv("NXO_TLS_VERIFY"); ok && (strings.ToLower(val) == "true") {
			c.TLSVerify = true
		}
		fmt.Printf("TLS Verify set to %v.\n", c.TLSVerify)
		//TLS Server and Client Config
		serverConfig, err := utils.GetTLServerConfig(c.ServerCert, c.ServerKey, c.CaCert, c.TLSVerify)
		fmt.Println("TLS server config created call returned.")
		if err != nil {
			fmt.Println("Error getting TLS server config, exiting...", err)
			log.Fatalf("Error getting TLS server config: %v", err)

		}
		fmt.Println("TLS server config created.")
		c.TLSServerConfig = serverConfig
		clientConfig, err := utils.GetTlsClientConfig(c.ClientCert, c.ClientKey, c.CaCert, c.TLSVerify)
		if err != nil {
			log.Fatalf("Error getting TLS client config: %v", err)
		}
		c.TlsClientConfig = clientConfig
		fmt.Println("TLS client config created.")
		fmt.Println("TLS configuration set.")
		//Facade / Proxy related info
		c.FacadeURL = utils.GetENV("NXO_FACADE_URL", "https://nxo-facade.example.com:8443")
		c.FacadeCred = utils.GetENV("NXO_FACADE_CRED", "admin:admin123")
		c.FacadeAuthType = utils.GetENV("NXO_FACADE_AUTH_TYPE", "Basic")

		c.EOProxyURL = utils.GetENV("NXO_EOPROXY_URL", "https://eoproxy.example.com:8443")
		c.EOProxyCred = utils.GetENV("NXO_EOPROXY_CRED", "admin:admin123")
		c.EOProxyAuthType = utils.GetENV("NXO_EOPROXY_AUTH_TYPE", "Basic")

		c.SaaSProxyAuthType = utils.GetENV("NXO_SAAS_PROXY_AUTH_TYPE", "Basic")
		c.SaaSProxyCred = utils.GetENV("NXO_SAAS_PROXY_CRED", "admin:admin123")
		c.SaaSProxyURL = utils.GetENV("NXO_SAAS_PROXY_URL", "https://saasproxy.example.com:8443")
		fmt.Println("Facade and Proxy configuration set.")
		//Server timeout properties
		tReadHrderTimeout := utils.GetENV("NXO_READ_HEADER_TIMEOUT", "10")
		if v1, err := strconv.Atoi(tReadHrderTimeout); err == nil {
			c.ReadHeaderTimeout = v1
		} else {
			c.ReadHeaderTimeout = 10
		}

		tReadTimeout := utils.GetENV("NXO_READ_TIMEOUT", "30")
		if v2, err := strconv.Atoi(tReadTimeout); err == nil {
			c.ReadTimeout = v2
		} else {
			c.ReadTimeout = 15
		}

		tWriteTimeout := utils.GetENV("NXO_WRITE_TIMEOUT", "30")
		if v3, err := strconv.Atoi(tWriteTimeout); err == nil {
			c.WriteTimeout = v3
		} else {
			c.WriteTimeout = 15
		}

		tIdleTimeout := utils.GetENV("NXO_IDLE_TIMEOUT", "60")
		if v4, err := strconv.Atoi(tIdleTimeout); err == nil {
			c.IdleTimeout = v4
		} else {
			c.IdleTimeout = 25
		}

		tHttpClientTimeout := utils.GetENV("NXO_HTTP_CLIENT_TIMEOUT", "10")
		if v5, err := strconv.Atoi(tHttpClientTimeout); err == nil {
			c.HttpClientTimeout = v5
		} else {
			c.HttpClientTimeout = 30
		}

		fmt.Println("Server timeout configuration set.")
		//Install Type onprem / SaaS
		c.InstallType = utils.GetENV("NXO_INSTALL_TYPE", "onprem")

		//Server Port
		c.PLUGINS_PORT = utils.GetENV("NXO_PLUGINS_PORT", "8443")
		fmt.Println("Configuration loaded successfully.")
		config = c
	})

	return config
}
