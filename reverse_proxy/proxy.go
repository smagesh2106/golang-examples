package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

var dynamicHeaders = map[string]string{} // store headers that may change
func init() {
	fmt.Println("Proxy with dynamic header refresh on failure")
}
func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <backend-url>", os.Args[0])
	}

	backendURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid backend URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // ⚠️ skips verification (use only for testing)
		},
	}
	// Director: inject headers dynamically
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		for k, v := range dynamicHeaders {
			req.Header.Set(k, v)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying %s %s", r.Method, r.URL.Path)
		tryProxy(w, r, proxy)
	})

	port := "9000"
	log.Printf("Starting proxy on :%s forwarding to %s", port, backendURL)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// tryProxy attempts proxy call with retry & header refresh
func tryProxy(w http.ResponseWriter, r *http.Request, proxy *httputil.ReverseProxy) {
	const maxRetries = 2
	//var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		rec := newResponseRecorder(w)
		proxy.ServeHTTP(rec, r)

		if rec.statusCode >= 200 && rec.statusCode < 300 {
			// success or client error, stop retrying
			rec.flush()
			return
		}

		// Save error
		//lastErr = rec.err

		// First failure → refresh headers
		if attempt == 0 {
			log.Println("First proxy attempt failed, fetching new headers...")
			refreshHeaders()
		}

		log.Printf("Retrying request (attempt %d)...", attempt+1)
		time.Sleep(1 * time.Second)
	}

	// If still failing
	//http.Error(w, "Proxy failed after retries: "+lastErr.Error(), http.StatusBadGateway)
	http.Error(w, "Proxy failed after retries: ", http.StatusBadGateway)
}

// refreshHeaders simulates calling external service to get new headers
func refreshHeaders() {
	resp, err := http.Get("http://localhost:8000/get-new-headers") // external service
	if err != nil {
		log.Printf("Failed to fetch new headers: %v", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// Example: expecting headers in format "Key:Value\nKey2:Value2"
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			dynamicHeaders[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	log.Printf("Updated headers: %+v", dynamicHeaders)
}

// --- Helper to capture proxy failures ---
type responseRecorder struct {
	w          http.ResponseWriter
	headers    http.Header
	statusCode int
	body       []byte
	err        error
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{w: w, headers: make(http.Header), statusCode: 0}
}

func (rr *responseRecorder) Header() http.Header {
	return rr.headers
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	rr.body = append(rr.body, b...)
	return len(b), nil
}

func (rr *responseRecorder) flush() {
	for k, vv := range rr.headers {
		for _, v := range vv {
			rr.w.Header().Add(k, v)
		}
	}
	if rr.statusCode != 0 {
		rr.w.WriteHeader(rr.statusCode)
	}
	if rr.body != nil {
		rr.w.Write(rr.body)
	}
}
