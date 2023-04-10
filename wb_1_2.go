package main

import (
	"fmt"
	"sync"
)

// вариант 1 используя waitGroup
// wg.Wait() блокирует фию пока все горутины не "сделают" Done
func wgsquare(inarr [5]int) {
	var wg sync.WaitGroup
	wg.Add(len(inarr))

	for _, n := range inarr {
		go func(wg *sync.WaitGroup, n int) {
			fmt.Printf("wg implementation %d \n", n*n)
			wg.Done()
		}(&wg, n)
	}

	wg.Wait()
}

// вариант 2 каналы
func chsquare(inarr [5]int) {
	ch := make(chan int)
	for _, n := range inarr {
		go func(n int) {
			fmt.Printf("ch implementation %d \n", n*n)
			ch <- n
		}(n)
	}
	// функция останется в этом цикле пока все элементы массива не передадутся в канал
	for range inarr {
		<-ch
	}
}

func main() {
	inarr := [5]int{2, 4, 6, 8, 10}
	wgsquare(inarr)
	chsquare(inarr)
}
