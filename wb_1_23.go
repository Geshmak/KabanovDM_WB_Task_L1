package main

import "fmt"

// верну пустой слайс и false в случае неудачи
func remove(sl []int, i int) ([]int, bool) {
	if i >= 0 && i <= len(sl)-1 {
		return append(sl[:i], sl[i+1:]...), true
	} else {
		var empty []int
		return empty, false
	}
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 7, 9, 10, 23, 8, 34}
	fmt.Println(sl)
	newsl, _ := remove(sl, 11)
	fmt.Println(newsl)
}
