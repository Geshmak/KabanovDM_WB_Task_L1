package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func powfoo(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
	fmt.Println("---powfoo завершился---")
	wg.Done()
}

func outfoo(ch2 <-chan int, wg *sync.WaitGroup) {
	for i := range ch2 {
		fmt.Println(i)
	}
	fmt.Println("---outfoo завершился---")
	wg.Done()
}

func main() {
	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)

	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	ch1 := make(chan int)
	ch2 := make(chan int)

	go powfoo(ch1, ch2, &wg)
	go outfoo(ch2, &wg)

	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] += rnd.Intn(50)
		ch1 <- arr[i]
	}
	close(ch1)

	fmt.Printf("Оригинальный массив: %v\n", arr)

}
