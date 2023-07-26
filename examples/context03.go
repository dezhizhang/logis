package examples

import (
	"context"
	"fmt"
	"time"
)

func worker1(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main() {

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Microsecond*50)
	wg.Add(1)

	go worker1(timeout)
	time.Sleep(time.Second * 5)
	cancelFunc()

	wg.Wait()
	fmt.Println("over")
}
