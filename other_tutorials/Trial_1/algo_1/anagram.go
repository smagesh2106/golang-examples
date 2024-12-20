package main

import (
	"fmt"
	//	"strconv"
	"time"
)

func main() {
	str1()
}

var res = [][]string{}
var arr = []string{"eat", "tea", "ate", "flow", "wolf", "pat", "tap", "bare", "bear", "spins", "olfw"}

func str1() {
	/*
		start := time.Now()
		for reduce2(arr) {
		}
		fmt.Println(time.Since(start))

		fmt.Println(res)
	*/

	start := time.Now()
	reduce3(arr)
	fmt.Println(time.Since(start))
	fmt.Println(res)

	/*
		m := make(map[string]int)
		v, ok := m["foo"]
		fmt.Println("%v\t%v\n", v, ok)
	*/
}

func reduce2(a []string) bool {
	if !(len(a) > 0) {
		return false
	}

	matched := []string{}
	nMatched := []string{}
	key := a[0]

	matched = append(matched, key)
	if len(a) > 1 {
		for i := 1; i < len(a); i++ {
			if anagram(key, a[i]) {
				matched = append(matched, a[i])
			} else {
				nMatched = append(nMatched, a[i])
			}
		}
	}

	arr = nMatched
	res = append(res, matched)
	if len(arr) > 0 {
		return true
	} else {
		return false
	}
}

func reduce3(a []string) {

	if !(len(a) > 0) {
		return
	}

	matched := []string{}
	nMatched := []string{}
	//	key := a[0]

	//matched = append(matched, key)
	matched = append(matched, a[0])
	if len(a) > 1 {
		for i := 1; i < len(a); i++ {
			//if anagram(key, a[i]) {
			if anagram(a[0], a[i]) {
				matched = append(matched, a[i])
			} else {
				nMatched = append(nMatched, a[i])
			}
		}
	}

	res = append(res, matched)
	if len(nMatched) > 0 {
		reduce3(nMatched)
	} else {
		return
	}
}

func anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	//	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
	for i := 0; i < len(a); i = i + 1 {
		//		v1, v2 := a[i], b[i]
		//		m1[v1]++
		//		m2[v2]++
		//		v1, v2 := a[i], b[i]
		m1[a[i]]++
		m2[b[i]]++

	}

	for k, _ := range m1 {
		v1, ok1 := m1[k]
		v2, ok2 := m2[k]
		if !(ok1 && ok2 && v1 == v2) {
			return false
		}
		/*
			if m1[k] != m2[k] {
				return false
			}
		*/
	}
	return true
}

func reduce(a *[]string) bool {
	fmt.Printf("------start-------%d\n", len(arr))
	for i := 0; i < len(*a); i++ {

	}

	*a = (*a)[1:]
	fmt.Printf("------end-------%d\n", len(arr))
	if len(*a) > 0 {
		return true
	} else {
		return false
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func bitCall(s string) uint8 {
	var r uint8 = 0
	for _, v := range s {

		if r == 0 {
			r = uint8(v)
			continue
		}

		r ^= uint8(v)
	}
	return r
}
