package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	data := ""
	for input.Scan() {
		data = strings.TrimSpace(input.Text())
		if data == "quit" {
			break
		}

		if len(data) > 1 {
			counts[data]++
		}
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
