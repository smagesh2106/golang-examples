package main

import (
	"fmt"
)

func main() {
	arr1()
	//cast()
	//str()
}

func cast() {
	f := 3.141
	fmt.Printf("f = %v\n", f)
	i := int(f)
	fmt.Printf("i = %v\n", i)
}

func str() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("c=%c\n", v)
	}
	s1 := []rune{}
	//	s1 = append(s1, e3)
	s1 = append(s1, 0xe383)
	s1 = append(s1, 0x97e3)
	s1 = append(s1, 0x83ad)
	s1 = append(s1, 0xe382)
	s1 = append(s1, 0xb0e3)
	s1 = append(s1, 0x83a9)
	//s1 = append(s1, e3)

	fmt.Println(s1)
	fmt.Println(string(s1))

	b := []byte(s)
	fmt.Println(b)

	r := [...]int{4: 0}
	fmt.Println(r)

}
func arr1() {
	fmt.Println("-------Array---------") //size must be specified
	//a := [3]int{} //this is also good
	var a [3]int
	for i, _ := range a {
		a[i] = (i + 1) * 10
	}
	fmt.Printf("a : %v\n", a)
	modify(a)
	fmt.Printf("a : %v\n", a)
	b := [...]int{20, 30, 40}
	fmt.Printf("\nb : %v\n", b)
	modify(b)
	fmt.Printf("b : %v\n", b)
	modify3(&b)
	fmt.Printf("b : %v\n", b)
	fmt.Println("\n\n-------Slice---------")
	s := []int{}
	for i := 0; i < 3; i++ {
		s = append(s, (i+1)*10)
	}
	fmt.Printf("\ns : %v\n", s)
	modify2(s)
	fmt.Printf("s : %v\n", s)

	t := make([]int, 0, 0)
	for i := 0; i < 3; i++ {
		t = append(t, (i+1)*10)
	}
	fmt.Printf("\nt : %v\n", t)
	modify2(t)
	fmt.Printf("t : %v\n", t)
	fmt.Printf("t=%v, len=%v, cap=%v\n", t, len(t), cap(t))
}

func modify(a [3]int) {
	a[0] = a[0] + 1
}

func modify2(a []int) {
	a[0] = a[0] + 1
}

func modify3(a *[3]int) {
	a[0] = a[0] + 1
}
