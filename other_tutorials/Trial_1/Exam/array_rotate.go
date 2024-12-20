package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swapArr(arr []int) {
	fmt.Printf("arr prt :%p\t %v\n", arr, arr)
	for i, j := 0, len(arr)-1; i <= (len(arr)-1)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]

	}

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
	fmt.Printf("nums prt :%p\t %v\n", nums, nums)
	swapArr(nums)
	fmt.Println(nums)
	//fmt.Println(ops)
}
