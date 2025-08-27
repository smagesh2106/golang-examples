package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// makeRequest executes the HTTP GET call with timeout
func makeRequest(url string, i int) (result string, err error) {
	// Recover from panic inside this function
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	client := http.Client{
		Timeout: 3 * time.Second, // request timeout
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if i < 4 {
		panic("simulated panic after reading body") // simulate a panic
	}
	return string(body), nil
}

// withRetry wraps the request with retry logic
func withRetry(url string, attempts int, delay time.Duration) (string, error) {
	var result string
	var err error

	for i := 1; i <= attempts; i++ {
		result, err = makeRequest(url, i)
		if err == nil {
			return result, nil
		}

		fmt.Printf("Attempt %d failed: %v\n", i, err)

		if i < attempts {
			time.Sleep(delay)
		}
	}
	return "", fmt.Errorf("all %d attempts failed: %v", attempts, err)
}

func main() {
	url := "https://httpbin.org/get" // test URL

	body, err := withRetry(url, 5, 1*time.Second)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Println("Response:", body)
}
