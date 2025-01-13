package main

import (
	"context"
	"fmt"
	"time"
)

func do4(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("do4 Error :%s\n", err.Error())
			}
			fmt.Println("do4 finished")
			return
		case num := <-ch:
			fmt.Printf("do4 received :%d\n", num)

		}
	}
}

func do3(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan int)
	//defer cancel()
	go do4(ctx, ch)
	for num := 1; num < 10; num++ {
		ch <- num
	}
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("do3 Finished")
}

func main() {
	ctx := context.Background()
	do3(ctx)

}
