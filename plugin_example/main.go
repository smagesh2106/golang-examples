package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"gopkg.in/yaml.v3"
)

// Define the interface
type Greeter interface {
	Greet(name string) string
}

// AgentPlugin represents a single plugin.

type AgentPlugin struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
}

// Config represents the top-level structure that contains AgentPlugins.
type Config struct {
	AgentPlugins []AgentPlugin `yaml:"AgentPlugins"`
}

func main() {

	// Open / Read the YAML file
	file, err := os.Open("plugins.yaml") // Adjust the file name as needed
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer file.Close()

	c := Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Path to the plugin directory
	pluginDir := "./plugins"

	// Open the plugin directory
	files, err := os.ReadDir(pluginDir)
	if err != nil {
		fmt.Println("Error reading plugin directory:", err)
		return
	}

	// Iterate over all files in the plugin directory
	for _, file := range files {
		fmt.Println("Found file:", file.Name())
		// Only process .so files
		if filepath.Ext(file.Name()) == ".so" {
			// Load the plugin
			pluginPath := filepath.Join(pluginDir, file.Name())
			p, err := plugin.Open(pluginPath)
			if err != nil {
				fmt.Printf("Error loading plugin %s: %v\n", file.Name(), err)
				continue
			}

			// Lookup for the Greeter implementation in the plugin
			var greeterSymbol interface{}
			greeterSymbol, err = p.Lookup(getName(c.AgentPlugins, file.Name()))
			if err != nil {
				fmt.Printf("Error looking up GreeterPlugin in %s: %v\n", file.Name(), err)
				continue
			}

			// Assert the loaded symbol to the Greeter interface
			greeter, ok := greeterSymbol.(Greeter)
			if !ok {
				fmt.Printf("%s does not implement the Greeter interface\n", file.Name())
				continue
			}

			// Call the Greet method in the plugin
			fmt.Printf("%s: %s\n", file.Name(), greeter.Greet("Alice"))
		}
	}
}

func getName(plugins []AgentPlugin, plugin string) string {

	retVal := ""
	for _, p := range plugins {
		//fmt.Println("======>", p.Path, plugin)
		if strings.Contains(p.Path, plugin) {
			retVal = p.Name
			break
		}
	}
	return retVal
}
