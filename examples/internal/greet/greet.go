package greet

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Welcome %s", name)
}
