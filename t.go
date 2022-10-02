package main

import (
	"fmt"
)

func firstGo(ch1 chan int, ch2 chan int) {
	for {
		select {
		case a, ok := <-ch1:
			if !ok {
				fmt.Print("first worker:", a)
				ch2 <- a
			}
		}
	}
}

func secondGo(ch2 chan int) {
	for {
		select {
		case a, ok := <-ch2:
			if !ok {
				fmt.Print("e worker:", a)
				ch2 <- a
			}
		}
	}
}

func main() {
	count := 2
	ch := make(chan int)
	ch2 := make(chan int)
	read := 0
	fmt.Scanln(&read)
	for i := 0; i < count; i++ {
		go firstGo(ch, ch2)
		go
	}
}
