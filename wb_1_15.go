package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

func createHugeString(size int) string {
	sur := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(sur)

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(rnd.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func SomeFunc() {
	v := createHugeString(1 << 10)

	//так нехорошо можно порезать руну!!!!!!!!!!!!!!!

	// justString = v[:100]
	justString = string(append([]rune{}, []rune(v)[:100]...)) // а так хорошо

	fmt.Println(justString)

}

func main() {
	SomeFunc()
}
