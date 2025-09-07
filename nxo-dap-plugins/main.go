package main

import (
	srv "golang-examples/nxo-dap-plugins/plugins/nxo-plugin/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start service in goroutine
	service := srv.GetNewNxoService()

	if err := service.Init(); err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	go func() {
		if err := service.Start(); err != nil {
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
