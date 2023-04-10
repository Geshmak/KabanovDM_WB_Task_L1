package main

import "fmt"

func reverse(str string) string {
	runes := []rune(str)
	first, last := 0, len(runes)-1

	for first < last {
		runes[first], runes[last] = runes[last], runes[first]
		first++
		last--
	}
	return string(runes)

}

func main() {
	str := "123456789"
	fmt.Println(reverse(str))
}
