package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingF      Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g⁰C", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%g⁰F", f) }
