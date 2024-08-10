package main

import (
	"context"
	"fmt"
)

func do1(ctx context.Context) {
	fmt.Printf("do something key/value,  key:%s\n", ctx.Value("my-key"))
}
func do2(ctx context.Context) {
	fmt.Printf("do something key/value,  key:%s\n", ctx.Value("my-key2"))
}

func main() {
	//Both context.Background() and context.TODO() creates empty context.
	// If you’re unsure which one to use, context.Background is a good default option

	ctx := context.Background()
	//ctx := context.TODO()
	ctx = context.WithValue(ctx, "my-key", "my-value")
	do1(ctx)
	ctx = context.WithValue(ctx, "my-key2", "my-value2")
	do2(ctx)
}
