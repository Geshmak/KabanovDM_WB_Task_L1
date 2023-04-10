package main

import (
	"fmt"
	"time"
)

// Sleep - pause a thread on time t
func Sleep(t time.Duration) {
	timer := time.NewTimer(t)

	<-timer.C
}

func main() {
	start := time.Now()
	fmt.Println("Начало")

	fmt.Println("Отсчет пошел")
	Sleep(time.Second * 2)

	fmt.Println("Время вышло ", time.Since(start))
}
