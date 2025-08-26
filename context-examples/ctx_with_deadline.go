package main

import (
	"context"
	"fmt"
	"time"
)

func do6(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("do6 Error :%s\n", err.Error())
			}
			fmt.Println("do6 finished")
			return
		case num := <-ch:
			fmt.Printf("do6 received :%d\n", num)
			time.Sleep(1 * time.Second)

		}
	}
}

func do5(ctx context.Context) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	ch := make(chan int)
	go do6(ctx, ch)

loop:
	for num := 1; num <= 5; num++ {
		select {
		case ch <- num:
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("do5 context done")
			break loop
		}
	}
	//cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("do5 Finished")
}

func main() {
	ctx := context.Background()
	do5(ctx)
}
