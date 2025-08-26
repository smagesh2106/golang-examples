// dbus_server.go
package main

import (
	"fmt"
	"log"

	dbus "github.com/godbus/dbus/v5"
	introspect "github.com/godbus/dbus/v5/introspect"
)

const serviceName = "com.example.GolangDBus"
const objectPath = "/com/example/GolangDBusObject"
const interfaceName = "com.example.GolangDBus"

// HelloService exposes methods on DBus
type HelloService struct{}

// Hello method implementation
func (s *HelloService) Hello() (string, *dbus.Error) {
	return "Hello <<<>>>> from Go D-Bus service!", nil
}

func main() {
	// Connect to the session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal("Failed to connect to session bus:", err)
	}
	defer conn.Close()

	// Request a well-known name
	reply, err := conn.RequestName(serviceName, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal("Failed to request name:", err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Fatal("Name already taken")
	}

	// Export the service (object + interface)
	hello := &HelloService{}
	conn.Export(hello, objectPath, interfaceName)

	// Export introspection data (so clients can query it)
	node := &introspect.Node{
		Name: string(objectPath),
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			{
				Name: interfaceName,
				Methods: []introspect.Method{
					{Name: "Hello", Args: []introspect.Arg{
						{Name: "response", Type: "s", Direction: "out"},
					}},
				},
			},
		},
	}
	conn.Export(introspect.NewIntrospectable(node), objectPath,
		"org.freedesktop.DBus.Introspectable")

	fmt.Println("✅ D-Bus service running... Press Ctrl+C to exit")
	select {} // block forever
}
