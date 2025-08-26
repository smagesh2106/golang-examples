// dbus_client.go
package main

import (
	"fmt"
	"log"
	"os"

	dbus "github.com/godbus/dbus/v5"
)

const serviceName2 = "com.example.GolangDBus"
const objectPath2 = "/com/example/GolangDBusObject"
const interfaceName2 = "com.example.GolangDBus"

func main() {
	// Connect to the session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatalf("Failed to connect to the DBus: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Call the Hello method
	call := conn.Object(serviceName2, objectPath2).Call(interfaceName2+".Hello", 0)

	// Check for errors
	if call.Err != nil {
		log.Fatalf("Failed to call method: %v", call.Err)
	}

	// Print the response
	var response string
	if err := call.Store(&response); err != nil {
		log.Fatalf("Failed to store response: %v", err)
	}

	fmt.Println("📩 Received response:", response)
}
