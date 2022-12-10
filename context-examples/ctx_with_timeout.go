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
			time.Sleep(100 * time.Millisecond)

		}
	}
}

func do5(ctx context.Context) {
	//deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	ch := make(chan int)
	go do6(ctx, ch)
	for num := 1; num <= 10; num++ {
		select {
		case ch <- num:
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			break
		}
	}
	//cancel()
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("do5 Finished")
}

func main() {
	ctx := context.Background()
	do5(ctx)

}
