package main

import (
	"fmt"
)

type Rectangle struct {
	length float64
	breath float64
}

func main() {
	var rect1 Rectangle
	rect1.length = 10.5
	rect1.breath = 5.5

	rect2 := Rectangle{}
	rect2.length = 11.5
	rect2.breath = 6.5

	rect3 := Rectangle{length: 12.1, breath: 9.66}
	var rect4 = []Rectangle{
		{1, 2},
		{3, 4},
		{4, 5},
	}

	//array of rectangles
	var rect5 [2]Rectangle

	rect5[0] = Rectangle{length: 1, breath: 2}
	rect5[1] = Rectangle{length: 2, breath: 1}

	//slice of rectangles
	var rect6 []Rectangle

	rect6 = append(rect6, Rectangle{length: 1, breath: 2})
	rect6 = append(rect6, Rectangle{length: 3, breath: 5})

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)
	fmt.Println(rect5)
	fmt.Println(rect6)
}
