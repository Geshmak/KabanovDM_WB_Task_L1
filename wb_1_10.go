package main

import "fmt"

//-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.

func main() {

	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -2.3, 22.4, -34.3, 2.3, -16.2, -4.1, 23.7}

	tmap := make(map[int][]float32)

	for _, fl := range arr {
		key := int(fl) - int(fl)%10
		if fl < 0 {
			key -= 10
		}
		tmap[key] = append(tmap[key], fl)
	}
	fmt.Println(tmap)
}
