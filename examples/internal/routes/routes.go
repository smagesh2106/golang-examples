package routes

import (
	"fmt"

	"github.com/smagesh2106/examples/internal/greet"
)

func Bye(name string) string {
	return fmt.Sprintf("Bye %s", greet.SayHello("Paul"))
}
