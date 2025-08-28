package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <backend-url>", os.Args[0])
	}

	backendURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid backend URL: %v", err)
	}

	// Create proxy with custom director
	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// Call the default director to set scheme/host etc.
		originalDirector(req)

		// Example: Modify headers depending on request type
		switch req.Method {
		case http.MethodPost, http.MethodPut:
			// Add custom header for write operations
			req.Header.Set("X-Proxy-Write", "true")
		case http.MethodGet:
			// Add tracing header for reads
			req.Header.Set("X-Proxy-Trace", "read-path:"+req.URL.Path)
		}

		// Example: Remove sensitive headers
		req.Header.Del("X-Unwanted-Header")

		// Example: Add auth header only for Redfish paths
		if strings.HasPrefix(req.URL.Path, "/redfish") {
			req.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQ=") // Base64 admin:password
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Proxying %s %s", r.Method, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	port := "9000"
	log.Printf("Starting proxy on :%s forwarding to %s", port, backendURL)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
