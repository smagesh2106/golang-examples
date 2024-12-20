package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		txt := input.Text()
		err := input.Err()
		if err != nil {
			fmt.Println(input.Err())
		}
		if txt == "quit" {
			break
		}
		counts[txt]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("Could not open file :%s\n", arg)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fs, err := os.Stat(arg)
			fs.Si
			b, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Printf("Could not read file :%v\n")
			}

			for _, str := range strings.Split(string(b), "\n") {
				counts[str]++
			}
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	//dup1()
	//dup2()
	dup3()
	//a := []int{1, 2, 3}
	//fmt.Println(a)
	//Array(a)
	//fmt.Println(a)
}

func Array(a []int) {
	a[0] = 9
	fmt.Printf("Inside func %v\n", a)
}

/*
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}
func main() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}
*/
