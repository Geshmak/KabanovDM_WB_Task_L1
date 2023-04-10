package main

import (
	"fmt"
	"math/rand"
	"time"
)

// быстрая сортиврока O(nlog n) обменов, лучше например чем вставка O(n^2):

// надо найти опорный элемент (у меня последний, но можно, вообще говоря, взять любой например центральный)
// все элементы меньше опорного положить слева от него, а больше справа

// повторяем то же самое, обрабатывая "отрезки" слева и справа от опорного рекурсивно! (остановимся когда длина подмассива равна 1)
func sort(m []int, first int, last int) {
	if first >= last {
		return
	} else {
		i := first
		j := first - 1
		pivot := last

		for i < pivot {
			if m[i] < m[pivot] {
				j++
				m[i], m[j] = m[j], m[i]

			}
			i++
		}
		j++
		m[i], m[j] = m[j], m[i]

		sort(m, first, j-1)
		sort(m, j+1, last)
	}
}

// Для удобства вызываю саму сортировку через эту фиию
func quicksort(m []int) []int {
	newm := make([]int, len(m))
	copy(newm, m)
	sort(newm, 0, len(m)-1)
	return newm
}

func randlist(n int) []int {
	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rnd.Intn(500)
	}
	return res
}
func main() {
	//m := []int{2, 8, 4, 7, 1, 3, 9, 6, 5}
	m := randlist(20)
	fmt.Println(m)
	newm := quicksort(m)

	fmt.Println(newm)
}
