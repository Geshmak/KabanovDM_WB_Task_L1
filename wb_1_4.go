package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(wnumber int, ch <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("---Воркер %d остановился---\n", wnumber)
			return
		case data := <-ch:
			fmt.Printf("Воркер %d прочитал: %v\n", wnumber, data)
			// do something
		}
	}
}

func main() {
	var wcount int
	fmt.Println("Введите кол-во Воркеров")
	fmt.Scanf("%d", &wcount)

	var wg sync.WaitGroup
	wg.Add(wcount)

	ch := make(chan int, wcount)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)

	for i := 0; i < wcount; i++ {
		go worker(i, ch, ctx, &wg)
	}

	for {
		select {
		case <-c:
			fmt.Printf("\n---Сигнал пойман---\n\n")
			cancel()
			wg.Wait()
			fmt.Println("---Программа завершена---")
			return
		default:

			ch <- rnd.Intn(100)
			time.Sleep(50 * time.Millisecond)
		}
	}

}
