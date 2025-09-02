package main

import (
	"fmt"
	"golang-examples/nxo-dap-config/config"
)

func main() {

	c := config.GetConfig()
	fmt.Println("Config loaded:", c.ONE, c.TWO, c.THREE)
	c.OUTPUT()
}
