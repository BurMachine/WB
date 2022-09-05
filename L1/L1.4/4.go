package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type workers struct {
	Channel chan interface{}
}

func (w *workers) work(id int) {
	log.Println(id, "Ready")
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case v := <-w.Channel:
			log.Println(v, "worker", id)
		}
	}
}

func main() {
	count := 0
	fmt.Println("Set worker's count:")
	fmt.Scanln(&count)
	fmt.Println("Workers:\n")
	ch := make(chan interface{}, 10)
	n := make([]workers, count)
	for i := 0; i < count; i++ {
		n[i] = workers{Channel: ch}
	}
	for i := 0; i < count; i++ {
		go n[i].work(i)
	}
	tic := time.Tick(1000 * time.Millisecond)
	for _ = range tic {
		ch <- rand.Intn(12)
	}
	time.Sleep(1 * time.Second)
}
