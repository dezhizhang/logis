package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ch <-chan bool) {
LABEL:
	for {
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("worker")
			time.Sleep(time.Second)

		}

	}
	wg.Done()
}

func main() {

	var ch = make(chan bool)
	wg.Add(1)

	go worker(ch)
	time.Sleep(time.Second * 5)
	ch <- true

	wg.Wait()

	fmt.Println("over")
}
