package main

import (
	"fmt"
	"math/rand"
	"time"
)

// создаю случайные мн-ва
var (
	sur rand.Source
	rnd *rand.Rand
)

func init() {
	sur = rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(sur)
}

func randlist(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rnd.Intn(20)
	}
	return res
}

// максимальный элемент в 2 мн-вах
func max(m1 []int, m2 []int) (max int) {
	for _, i := range append(m1, m2...) {
		if i > max {
			max = i
		}
	}
	return
}

// пересечение
func intersect(a []int, b []int) (res []int) {
	bools := make([]bool, max(a, b)+1)

	//bools == true на индексах первого массива
	for _, v := range a {
		bools[v] = true
	}
	// в res записываются только те зн из 2го массива по индексам которых bools возвращает true
	for _, v := range b {
		if bools[v] {
			res = append(res, v)
			bools[v] = false // позволяет не писать дубли
		}
	}

	return
}

func main() {
	a := randlist(20)
	b := randlist(20)

	//a := []int{7, 15, 4, 6, 12}
	//b := []int{19, 15, 1, 15, 15}

	//a := []int{21, 7, 15, 12, 19, 6, 10, 11, 8, 14, 0, 3, 9, 5, 13}
	//b := []int{3, 17, 13, 6, 1, 21, 4, 19, 22, 14, 18, 7, 2, 16, 15, 8}
	fmt.Println(a, b)
	res := intersect(a, b)
	fmt.Println(res)
}
