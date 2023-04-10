package main

import (
	"fmt"
	"strings"
)

//Для каждой руны есть соответственное зн в мапе
// если значение true и руна снова встретилась вернуть false

func isUnique(str string) bool {
	str = strings.ToLower(str)
	boolmap := make(map[rune]bool)
	for _, rune := range str {
		_, ok := boolmap[rune]
		if ok {
			return false
		}
		boolmap[rune] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcHdeFg"))
	fmt.Println(isUnique("abcHdeFgh"))
	fmt.Println(isUnique("abcHdeFga"))
}
