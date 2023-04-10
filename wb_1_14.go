package main

import "fmt"

func main() {
	var slice []interface{}
	ch := make(chan int)
	slice = append(slice, 12, 5.5, "string", true, ch)

	for _, val := range slice {
		switch v := val.(type) {
		case int:
			fmt.Println("int:", v)
		case float64:
			fmt.Println("float64:", v)
		case string:
			fmt.Println("string:", v)
		case bool:
			fmt.Println("bool:", v)
		case chan int:
			fmt.Println("channel:", v)
		default:
			fmt.Println("unknown")
		}
	}
}
