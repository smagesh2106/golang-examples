package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//find immediate sums in an array
func calPoints(arr []int) int {
	var res int = 0

	for i := 0; i <= len(arr)-2; i = i + 1 {
		for j := 1; j <= len(arr)-1; j = j + 1 {
			//			fmt.Printf("i=%d, arr[i]=%d,    j=%d, arr[j]=%d\n", i, arr[i], j, arr[j])
			if arr[i]+1 == arr[j] {
				//				fmt.Println("found")
				res += 1
				break
			}
		}
	}

	//	fmt.Println("arr = ", arr)
	return res
}

func calPoints2(arr []int) int {
	var res int
	found := make(map[int]bool)
	// code here
	for _, val := range arr {
		found[val] = true
	}

	for _, val := range arr {
		if found[val+1] {
			res++
		}
	}

	return res
}

func main() {

	nums := []int{}
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')

	rawInput = strings.Replace(rawInput, "\n", " ", -1)

	ops := strings.Split(rawInput, " ")

	for index := range ops {
		ele, err := strconv.Atoi(ops[index])
		if err == nil {
			nums = append(nums, ele)
		}
	}
	fmt.Println("---old-->", calPoints(nums))
	fmt.Println("---new-->", calPoints2(nums))
	//fmt.Println(ops)
}
