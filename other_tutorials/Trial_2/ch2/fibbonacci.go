package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("Fib of %v  =  %v\n", 75, fib(75))
	//fmt.Printf("Gcd of %v, %v  =  %v\n", 20, 25, gcd(20, 25))
	lcm(52, 84)
}

func fib(n int) uint64 {
	var x, y uint64 = 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Printf("%v\t%v\n", x, y)
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		fmt.Printf("Before -> X=%v,   Y=%v\n", x, y)
		x, y = y, x%y
		fmt.Printf("After  -> X=%v,   Y=%v\n\n", x, y)
	}

	return x
}

func lcm(x, y int) {
	m, n := make(map[int]int), make(map[int]int)
	mA, nA := []int{2, 3}, []int{2, 3}
	num, flag := 1, false

	for i, j := 4, 4; i < x/2 || j < y/2; i, j = i+1, j+1 {
		flag = false

		if i == j || i > j {
			num = i
		} else {
			num = j
		}
		if CheckPrime(num) {
			flag = true
		}
		if i < x/2 && flag {
			mA = append(mA, num)
		}
		if j < y/2 && flag {
			nA = append(nA, num)
		}
	}

	fmt.Printf("mA = %v\n", mA)
	fmt.Printf("nA = %v\n", nA)

	notFound := false
	for notFound == false {
		flag := false
		for _, v := range mA {
			if x%v == 0 {
				fmt.Printf("x=%d, \tv=%d\n", x, v)
				m[v]++
				x, flag = x/v, true
				break
			}
		}
		if flag == false {
			notFound = true
		}
	}
	if len(m) == 0 {
		m[x] = 1
	}

	fmt.Println()
	notFound = false
	for notFound == false {
		flag := false
		for _, v := range nA {
			if y%v == 0 {
				fmt.Printf("y=%d, \tv=%d\n", y, v)
				n[v]++
				y, flag = y/v, true
				break
			}
		}
		if flag == false {
			notFound = true
		}
	}

	if len(n) == 0 {
		n[y] = 1
	}

	fmt.Printf("\nm=%v\nn=%v\n", m, n)
	res := make(map[int]int)
	for key, val1 := range m {

		if val2, ok := n[key]; ok {
			if val2 > val1 {
				val1 = val2
			}
		}
		res[key] = val1
	}
	for key, val := range n {
		res[key] = val
	}

	value := 1
	for key, val := range res {
		value = value * key * val
	}
	fmt.Printf("res = %v\t\t, value = %v\n", res, value)
}

func CheckPrime(number int) bool {
	isPrime := true
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime == true {
		return true
	} else {
		return false
	}
}

//-------------
/*
func lcm(x, y int) {
	m := make(map[int]int)
	n := make(map[int]int)

	//mA := make([]int, x/2)
	//nA := make([]int, y/2)
	mA := []int{2, 3}
	nA := []int{2, 3}

	//for i := 0; i < x/2; i++ {
	for i := 4; i < x/2; i++ {
		//mA[i] = i + 1
		if CheckPrime(i) {
			mA = append(mA, i)
		}
	}
	//for i := 0; i < y/2; i++ {
	for i := 4; i < y/2; i++ {
		//nA[i] = i + 1
		if CheckPrime(i) {
			nA = append(nA, i)
		}
	}
	fmt.Printf("mA = %v\n", mA)
	notFound := false
	for notFound == false {
		flag := false
		//for _, v := range mA[1:] {
		for _, v := range mA {
			if x%v == 0 {
				fmt.Printf("x=%d, \tv=%d\n", x, v)
				x = x / v
				m[v]++
				flag = true
				break
			}
		}
		if flag == false {
			notFound = true
		}
	}
	if len(m) == 0 {
		m[x] = 1
	}

	notFound = false
	for notFound == false {
		flag := false
		//for _, v := range nA[1:] {
		for _, v := range nA {
			if y%v == 0 {
				fmt.Printf("y=%d, \tv=%d\n", y, v)
				y = y / v
				n[v]++
				flag = true
				break
			}
		}
		if flag == false {
			notFound = true
		}
	}

	if len(n) == 0 {
		n[y] = 1
	}

	fmt.Printf("m=%v\n", m)
	fmt.Println("\n\n")
	fmt.Printf("n=%v\n", n)

}
*/
