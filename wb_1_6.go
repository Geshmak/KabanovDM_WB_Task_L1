package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// // остановка каналом
// нужно закрыть канал чтобы завершить горутину
func foo1(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		fmt.Println("foo1 working")
		time.Sleep(200 * time.Millisecond)
	}
	println("foo1 end")
	return

}

// остановка контекстом
// в main надо вызвать фию cancel()
func foo2(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo2 end")
			return
		default:
			fmt.Println("foo2 working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// // остановка каналом
// нужно написать в канал чтобы завершить горутину
func foo3(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			println("foo3 end")
			return
		default:
			fmt.Println("foo3 working")
			time.Sleep(200 * time.Millisecond)
		}
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//1
	go foo1(ch1, &wg)
	//2
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go foo2(ch2, &wg, ctx)

	//3
	go foo3(ch3, &wg)

	time.Sleep(1 * time.Second)

	close(ch1) //1 end
	cancel()   //2 end
	ch3 <- 1   //3 end

	wg.Wait()
}
