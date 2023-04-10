package main

import "fmt"

func chsquare(inarr [5]int) (res int) {
	ch := make(chan int)
	for _, n := range inarr {
		go func(n int) {
			ch <- n * n
		}(n)
	}

	for range inarr {
		res += <-ch
	}
	return
}

func main() {
	inarr := [5]int{2, 4, 6, 8, 10}
	fmt.Println(chsquare(inarr))
}
