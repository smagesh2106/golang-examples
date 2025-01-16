package main

import "fmt"

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) int {
	*p++
	return *p
}
func main4() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g⁰F = %g⁰C", freezingF, FtoC(freezingF))
	var names []string
	fmt.Println()
	fmt.Println(names)
	v := 1
	p := &v
	incr(p)
	fmt.Println(incr(p))
}
