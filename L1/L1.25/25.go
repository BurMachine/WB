package main

import (
	"fmt"
	"time"
)

// SecondSleep
// Реализация через тикер
func SecondSleep(duration time.Duration) time.Time {
	ticker := time.Tick(duration)
	for val := range ticker {
		return val
	}
	return time.Now()
}

// Sleep
//Реализация через канал time.After
func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("1")
	Sleep(3 * time.Second)
	//SecondSleep(3 * time.Second)
	fmt.Println("2")
	time.Sleep()
}
