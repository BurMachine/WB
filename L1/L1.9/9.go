package main

import (
	"fmt"
	"sync"
)

func channelProc(ch1, ch2 chan int) {
	for value := range ch1 {
		ch2 <- value * value
	}
	close(ch2)
}

func channelOut(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range channel {
		fmt.Println(value)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := &sync.WaitGroup{}
	go channelProc(ch1, ch2)
	wg.Add(1)
	go channelOut(ch2, wg)
	for _, val := range arr {
		ch1 <- val
	}
	close(ch1)
	wg.Wait()
}
