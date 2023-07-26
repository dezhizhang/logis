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
### WithDeadline
```go
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

```
### WithTimeout
```go
var wg sync.WaitGroup

func worker(ctx context.Context) {
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

	go worker(timeout)
	time.Sleep(time.Second * 5)
	cancelFunc()

	wg.Wait()
	fmt.Println("over")
}

```