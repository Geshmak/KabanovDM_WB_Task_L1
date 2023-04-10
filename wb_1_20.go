package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	var newstr string

	strslice := strings.Split(str, " ")
	for i := len(strslice) - 1; i >= 0; i-- {
		newstr += strslice[i] + " "
	}

	return newstr
}

func main() {

	str := "snow dog sun check split extra"
	fmt.Println(reverseWords(str))
}
