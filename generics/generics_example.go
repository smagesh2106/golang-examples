package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}

func main() {
	xs := []int{3, 5, 10}
	total := Sum(xs)
	fmt.Println(total)

	xm := []int64{3, 5, 10}
	total2 := Sum(xm)
	fmt.Println(total2)

	xl := []float64{5.0, 5.0, 10.5}
	total3 := Sum(xl)
	fmt.Println(total3)

}
