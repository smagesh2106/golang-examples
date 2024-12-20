package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	timeout := 10 * time.Second
	deadLine := time.Now().Add(timeout)
	url := "http://www.rrrdiff.com"
	for tries := 0; time.Now().Before(deadLine); tries++ {
		_, err := http.Head(url)
		if err == nil {
			fmt.Println("Done...")
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("...retrying....")
		if tries == 3 {
			url = "http://www.rediff.com"
		}

	}
	fmt.Println("Wakeup")
}
