package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-interruptCh:
			fmt.Println("\nInput interrupted!")
			for line, n := range count {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}

		os.Exit(0) // Exit the program when interrupted
	}()
	str := "hello"
	for _, c := range str {
		fmt.Printf("%c\n", c)
	}
	for input.Scan() {
		count[input.Text()]++
	}
	/*
		for line, n := range count {
			fmt.Printf("%d\t%s\n", n, line)
		}
	*/
}
