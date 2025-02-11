// dbus_server.go
package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

const serviceName = "org.freedesktop.DBus2"
const objectPath = "/org/freedesktop/DBus2"
const interfaceName = "org.freedesktop.DBus2"

type HelloService struct{}

func (h *HelloService) Hello() string {
	return "Hello from DBus Server!"
}

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Println("Error listening for DBus name:", err)
		panic(err)
	}
	defer conn.Close()

	reply, err := conn.RequestName(serviceName, dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Println("Name already taken")
		return
	}
	// Register object path with Hello method
	conn.Export(new(HelloService), objectPath, interfaceName)

	// Wait for incoming DBus requests
	fmt.Println("Server running...")

	select {}
}
