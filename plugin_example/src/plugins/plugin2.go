package main

import (
	"fmt"
)

type GreeterPlugin2 struct{}

// Greeter implementation for plugin2

// Implement the Greet method
func (g GreeterPlugin2) Greet(name string) string {
	return fmt.Sprintf("Greetings from Plugin2, %s!", name)
}

// Export GreeterPlugin to be used by main.go
var GreeterPluginInstance2 GreeterPlugin2
