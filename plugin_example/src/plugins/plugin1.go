package main

import "fmt"

// Greeter implementation for plugin1
type GreeterPlugin1 struct{}

// Implement the Greet method
func (g GreeterPlugin1) Greet(name string) string {
	return fmt.Sprintf("Hello from Plugin1, %s!", name)
}

// Export GreeterPlugin to be used by main.go
var GreeterPluginInstance1 GreeterPlugin1
