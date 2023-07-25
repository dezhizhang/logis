package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}
