package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(n *counter, wg *sync.WaitGroup, stop <-chan int, name int) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Printf("---foo%d stopped---\n", name)
			return
		default:
			fmt.Printf("foo%d working\n", name)

			n.Lock()
			n.num++
			n.Unlock()

			time.Sleep(time.Millisecond * 200)
			fmt.Printf("foo%d resting\n", name)
		}

	}
}

type counter struct {
	num int
	sync.Mutex
}

func main() {
	var (
		wg sync.WaitGroup
	)
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)

	n := counter{num: 0}

	go foo(&n, &wg, ch1, 1)
	go foo(&n, &wg, ch2, 2)

	time.Sleep(2 * time.Second)

	ch1 <- 1
	ch2 <- 1
	fmt.Println(n.num)
	wg.Wait()
}
