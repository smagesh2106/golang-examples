package main

import (
	"crypto/sha256"
	"fmt"
)

func main22() {
	//var a [3]int
	//var b [4]int = [4]int{1, 2, 3, 4}
	var c = [...]int{4: -1}
	fmt.Println("len of c: ", len(c))
	for i, v := range c {
		fmt.Printf("%d\t%d\n", i, v)
	}
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n", c1, c2)
	var arr [32]byte

	zero(&arr)
	select {}

	var months = [...]string{1: "Jan", 2: "Feb", 3: "Mar"}

	for i, m := range months {
		fmt.Printf("-->%d\t%s\n", i, m)
	}

	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	reverse(arr1[:])
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	reverse2(&arr2)
	fmt.Println(arr2)

	var arr3 = []int{1, 2, 3, 4, 5}
	arr3 = append(arr3, 6, 7, 8)

	fmt.Println(arr3)
	reverse3(&arr3)
	fmt.Println(arr3)

}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// pass array
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// pass ptr to an array
func reverse2(ptr *[5]int) {
	left, right := 0, len(ptr)-1
	for left < right {
		ptr[left], ptr[right] = ptr[right], ptr[left]
		left++
		right--
	}
}

// pass ptr to an slice
func reverse3(ptr *[]int) {
	left, right := 0, len(*ptr)-1
	for left < right {
		(*ptr)[left], (*ptr)[right] = (*ptr)[right], (*ptr)[left]
		left++
		right--
	}
}
