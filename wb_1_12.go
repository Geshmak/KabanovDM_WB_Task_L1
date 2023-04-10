package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	boolmap := make(map[string]bool)

	for _, v := range arr {
		boolmap[v] = true
	}
	for key := range boolmap {
		res = append(res, key)
	}

	fmt.Println(res)
}
