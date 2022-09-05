package main

import (
	"fmt"
	"time"
)

func say(word string) {
	fmt.Println(word)
}

func main() {
	ch := make(chan int)
	go sayHello(ch)
	fmt.Println("kuku")
	for i := range ch {
		fmt.Println(i)
	}
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit)
}
