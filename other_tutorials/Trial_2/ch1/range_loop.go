package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Max(x, y int) float64 {
	return math.Max(float64(x), float64(y))
}

func Range() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "-"
	}
	fmt.Println(s)
	s, sep = "", ""
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		s += sep + arg
		sep = "-"
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], "+"))

	//range oper on map
	m := map[string]string{
		"dog":  "bones",
		"cat":  "fish",
		"bird": "grains",
	}
	for key, val := range m {
		fmt.Printf("%s\t%s\n", key, val)
	}
	ChangeMap(m)
	fmt.Println("\n\n")
	for key, val := range m {
		fmt.Printf("%s\t%s\n", key, val)
	}

}

func main() {
	fmt.Println(Max(1, 2))
	Range()
}

func ChangeMap(m map[string]string) {
	m["dog"] = "meat"
}
