package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func product(arr []int, index int) int64 {
	var res int64 = 1
	fmt.Println(arr, index)
	for i, v := range arr {
		if i != index-1 {
			res = res * int64(v)
		}
	}
	return res
}

func main() {

	nums := []int{}
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	rawInput = strings.Replace(rawInput, "\n", " ", -1)

	var index int
	fmt.Scanf("%d", &index)

	ops := strings.Split(rawInput, " ")

	for index := range ops {
		ele, err := strconv.Atoi(ops[index])
		if err == nil {
			nums = append(nums, ele)
		}
	}

	fmt.Println(product(nums, index))
	//fmt.Println(ops)
}
