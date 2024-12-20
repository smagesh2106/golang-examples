package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calPoints(ops []string) int {
	var res int = 0
	var values = []int{}
	for _, ele := range ops {
		switch ele {
		case "C":
			fmt.Println("C :", ele)
			if len(values) >= 1 {
				values = values[:len(values)-1]
			}
		case "D":
			fmt.Println("D :", ele)
			if len(values) >= 1 {
				values = append(values, values[len(values)-1]*2)
			}
		case "+":
			fmt.Println("+ :", ele)
			if len(values) >= 2 {
				num := values[len(values)-1] + values[len(values)-2]
				values = append(values, num)

			}
		default:
			fmt.Println("Default :", ele)
			i, err := strconv.Atoi(ele)
			if err == nil {
				values = append(values, i)
			}
			fmt.Println(values)
		}

	}
	for _, val := range values {
		res += val
	}
	return res
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')

	rawInput = strings.Replace(rawInput, "\n", " ", -1)

	ops := strings.Split(rawInput, " ")

	fmt.Println(calPoints(ops))
	//fmt.Println(ops)
}
