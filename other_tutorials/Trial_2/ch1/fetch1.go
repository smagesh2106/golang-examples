package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var lock sync.Mutex
var count int

func Fetch1() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Could not fetch url :%s\n", url)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Could not read the body")
		}
		fmt.Printf("%s\n", b)
	}
}

func Fetch2(url string, ch chan<- string) {
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("Could not fetch url :%s\n", url)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Could not read the body :%v", err.Error())
		return
	}
	ch <- string(b) + "\n\n========================== END END END ==========================\n\n"
	lock.Lock()
	count++
	defer lock.Unlock()
}

func main() {
	//Fetch1()
	urls := os.Args[1:]
	ch := make(chan string, 10)
	for _, url := range urls {
		go Fetch2(url, ch)
	}
	/*
		select {
		case dat := <-ch:
			fmt.Println(dat)
		}
	*/

	for range urls {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		}
	}
	fmt.Printf("Number of URLs :%d\n", count)
}
