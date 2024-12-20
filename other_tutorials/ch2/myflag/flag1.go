package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omit Trailing new line")
var sep = flag.String("s", " ", "Seperator")

var medals = []string{"Gold", "Silver", "Bronze"}

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !(*n) {
		fmt.Println()
	}

	fmt.Println(medals)
}
