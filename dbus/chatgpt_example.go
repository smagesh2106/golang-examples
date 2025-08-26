package main

import (
	"fmt"
	"log"

	dbus "github.com/godbus/dbus/v5"
)

func main() {
	// Connect to the session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal("Failed to connect to session bus:", err)
	}
	defer conn.Close()

	// ----------------------------------------------------
	// 1. Method Call: Call ListNames() on the D-Bus daemon
	// ----------------------------------------------------
	obj := conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")

	var names []string
	err = obj.Call("org.freedesktop.DBus.ListNames", 0).Store(&names)
	if err != nil {
		log.Fatal("Method call failed:", err)
	}
	fmt.Println("📌 Registered bus names (apps/services):")
	for _, n := range names {
		fmt.Println("-------> :", n)
	}

	// ----------------------------------------------------
	// 2. Signal: Listen for NameOwnerChanged events
	// ----------------------------------------------------
	// Add a match rule to receive signals
	err = conn.AddMatchSignal(
		dbus.WithMatchInterface("org.freedesktop.DBus"),
		dbus.WithMatchMember("NameOwnerChanged"),
	)
	if err != nil {
		log.Fatal("Failed to add match for signals:", err)
	}

	// Make a channel to receive signals
	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	fmt.Println("\n📡 Listening for NameOwnerChanged signals (Ctrl+C to exit)...")

	// Run in background
	go func() {
		for sig := range c {
			fmt.Printf("Signal received: %s %v\n", sig.Name, sig.Body)
		}
	}()

	// ----------------------------------------------------
	// 3. Property: Get "Features" property from org.freedesktop.DBus
	// ----------------------------------------------------
	var features []string
	err = obj.Call("org.freedesktop.DBus.Get", 0,
		"org.freedesktop.DBus", "Features").Store(&features)
	if err != nil {
		fmt.Println("⚠️ Could not read Features property (may not be supported):", err)
	} else {
		fmt.Println("\n🔧 D-Bus daemon features:", features)
	}

	// Keep the program running to receive signals
	select {}
}
