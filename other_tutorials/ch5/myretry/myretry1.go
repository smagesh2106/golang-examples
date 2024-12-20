package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("unable to create log file.")
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
}

func main() {
	log.Println("startign log")
	/*
		fmt.Println("Commencing countdown. Press returnto abort.")
		select {
		case <-time.After(10 * time.Second): // Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	*/
	/*
		if e := WaitForServer("http://www.rrrediff.com"); e != nil {
			fmt.Printf("%v", e)
		}
	*/

	for {
		if e := WaitForServer("http://www.rediff.com"); e != nil {
			fmt.Printf("%v", e)
		} else {
			break
		}
	}

}

func WaitForServer(url string) error {
	const timeout = 15 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			fmt.Println(" No errors....")
			return nil // success
		}
		//fmt.Printf("server not responding (%s);retrying...\n", err)
		fmt.Printf("sleep for  (%v, %v, %v), retrying...\n", time.Second, uint(tries), time.Second<<uint(tries))
		time.Sleep(time.Second << uint(tries)) //exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
