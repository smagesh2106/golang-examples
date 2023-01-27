package main

import (
	"fmt"
	m "golang-examples/lib"
)

func main() {
	res := permutations("abc")

	for s := range res {
		fmt.Println(res[s])
	}
}

func recursivePermutations(str string, prefix string, result *[]string, callstack *m.Stack) {
	if len(str) == 0 {
		*result = append(*result, prefix)
		//callstack.Pop()
		callstack.Print()
		callstack.Pop()
	} else {
		for i := range str {
			rem := str[:i] + str[i+1:]
			//fmt.Printf("i=%d<--->%s--->%s<---->%s<---->%d,   %s+%s=%s\n", i, str[:i], str[i+1:], rem, len(rem), prefix, string(str[i]), prefix+string(str[i]))
			callstack.Push(fmt.Sprintf("i=%d, str=%s,  len(str)=%d, str[:i]=%s,  str[i+1:]=%s, str[i]=%s, rem=%s, prefix=%s,  prefix+str[i]=%s\n", i, str, len(str), str[:i], str[i+1:], string(str[i]), rem, prefix, prefix+string(str[i])))
			recursivePermutations(rem, prefix+string(str[i]), result, callstack)
		}
	}
}

func permutations(str string) []string {
	mystack := m.NewStack(25)
	result := []string{}
	recursivePermutations(str, "", &result, mystack)
	return result
}

//++++++++++++++++
