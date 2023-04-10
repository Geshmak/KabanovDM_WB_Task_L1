package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func foo(rwm *RWmap, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	i := 0
	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)

	for {
		select {
		case <-ch:
			fmt.Println("foo завершился")
			return
		default:
			key := rnd.Intn(5)
			rwm.mux.Lock()
			rwm.m[key] = i
			rwm.mux.Unlock()
			i++

			fmt.Printf("Положил %d по ключу %d\n", i, key)
		}
		time.Sleep(105 * time.Millisecond)
	}
}

func bar(rwm *RWmap, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)

	for {
		select {
		case <-ch:
			fmt.Println("bar завершился")
			return
		default:
			key := rnd.Intn(5)
			rwm.mux.RLock()
			val, ok := rwm.m[key]
			rwm.mux.RUnlock()
			if ok {
				fmt.Printf("Достал %d по ключу %d\n", val, key)
			} else {
				fmt.Printf("Искал по ключу %d, но не нашел ничего\n", key)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

type RWmap struct {
	m   map[int]int
	mux sync.RWMutex
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch1 := make(chan int)
	ch2 := make(chan int)

	rwm := RWmap{m: make(map[int]int)}

	// bar читает и спользует Rlock(), те блокируется если кто-то использует Lock()
	// foo пишет и использует Lock()
	go foo(&rwm, &wg, ch1)
	go bar(&rwm, &wg, ch2)
	// получается несколько читателей таких как bar могут получить доступ одновременно
	// но только один писатель типа foo получет доступ в один момент,
	// при этом читатели тоже но могут получить доступ в этот момет

	time.Sleep(3 * time.Second)

	ch1 <- 1
	ch2 <- 1

	wg.Wait()
}
