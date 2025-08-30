package main

import (
	"fmt"
	"golang-examples/nxo-dap-plugins/plugins/nxo-plugin/models"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// import "golang-examples/nxo-dap-plugins/plugins/nxo-plugin/service"
func main() {
	host := "0.0.0.0"
	port := "8443"
	fmt.Printf("Starting server at %s:%d\n", host, port)
	// Here you would typically start your server, e.g., http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)

	// Start service in goroutine
	service := models.GetNewNxoService("onprem", "admin", host, port)
	if err := service.Init(); err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	go func() {
		if err := service.Start(); err != nil && err.Error() != "http: Server closed" {
			log.Fatalf("Service error: %v", err)
		}
	}()

	// Wait for Ctrl-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan // blocks until user presses Ctrl-C

	log.Println("Shutting down...")
	if err := service.Stop(); err != nil {
		log.Printf("Error stopping service: %v", err)
	} else {
		log.Println("Service stopped gracefully")
	}
}
