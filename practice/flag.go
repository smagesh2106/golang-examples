package main

import (
	"flag"
	"fmt"
	"golang-examples/practice/tempconv"
	"strings"
	"time"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")
var j = flag.Bool("j", false, "say hello")

func init() {
	fmt.Println("...initializing.")
}

func HasPrefix(s, prefix string) bool {
	return ((len(s) >= len(prefix)) && (s[:len(prefix)] == prefix))
}

func HasSuffix(s, suffix string) bool {
	return ((len(s) >= len(suffix)) && (s[len(s)-len(suffix):] == suffix))
}

func Contains(s, substr string) bool {
	found := false
	for i := 0; i <= (len(s) - len(substr)); i++ {
		if !found && HasPrefix(s[i:], substr) {
			found = true
			break
		}
	}
	return found
}

func main6() {
	flag.Parse()
	if !*j {
		fmt.Println(strings.Join(flag.Args(), *sep) + ", say hello")
	} else {
		fmt.Println(strings.Join(flag.Args(), *sep))
	}
	if !*n {
		fmt.Println()
	}

	var x interface{}
	v, ok := x.(int)
	if ok {
		fmt.Println(v)
	}
	fmt.Println(tempconv.AbsoluteZeroC)
	st := "hello world"
	fmt.Println((len(st)))
	fmt.Println(st[0], st[7])
	fmt.Printf("%c, %c\n", st[0], st[3])
	fmt.Println(st[0:5])

	fmt.Printf("--->%t\n", HasPrefix("hello world", "hello"))
	fmt.Printf("--->%t\n", HasSuffix("hello world", "world"))
	fmt.Printf("--->%t\n", Contains("hello world", "rld"))

	for i, c := range st {
		fmt.Printf("%c", c)
		if i < len(st)-1 {
			fmt.Printf("->")
		}

	}
	fmt.Println()
	b := []byte(st)
	r := []rune(st)
	fmt.Println(b)
	fmt.Println(r)

	//type Weekday int
	const (
		//Sunday Weekday = iota
		Sunday int = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Monday)

	type Flags uint
	const (
		FlagUp Flags = 1 << iota // is up
		FlagBroadcast
		// supports broadcast access capability
		FlagLoopback
		// is a loopback interface
		FlagPointToPoint
		// belongs to a point-to-point link
		FlagMulticast
		// supports multicast access capability
		AA
		BB
		CC
		DD
		EE
	)

	fmt.Println(EE)

	var arr [100]float64
	var i uint64
	for i = 0; i < 100; i++ {
		ii := 1 << i
		arr[i] = float64(ii)
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d\t%f\n", i, arr[i])
	}

}
