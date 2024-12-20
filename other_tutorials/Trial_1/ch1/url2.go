package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)

	nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d %s", sec, nbytes, url)
}
