package main

import (
	"fmt"
	"sync"
	"time"
)

func readfoo(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		foo int
		ok  bool
	)
	for {
		foo, ok = <-ch
		if !ok {
			println("!Читатель остановился")
			return
		}
		fmt.Printf("Читатель прочитал %v\n", foo)

	}

}

func main() {
	var (
		dur   int
		wg    sync.WaitGroup
		count int
	)
	wg.Add(1)
	ch := make(chan int)

	fmt.Println("Введите кол-во секунд")
	fmt.Scanf("%d", &dur)

	var duration = time.Second * time.Duration(dur)
	timer := time.NewTimer(duration)
	fmt.Printf("Отсчет пошел (%d секунд)\n", dur)

	go readfoo(ch, &wg)

	for {
		select {
		case <-timer.C:
			fmt.Println("---Время вышло---")

			close(ch)
			wg.Wait()

			println("!Писатель остановился")
			return
		default:
			count++
			fmt.Printf("Писатель написал %v\n", count)
			ch <- count

			time.Sleep(200 * time.Millisecond)
		}
	}
}
