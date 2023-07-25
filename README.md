# logis

### context
```go
var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("working...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:

		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//var ch = make(chan bool)
	wg.Add(1)

	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()

	fmt.Println("over")
}

```
