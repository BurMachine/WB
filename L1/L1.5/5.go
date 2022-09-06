package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type worker struct {
	Channel chan interface{}
	timeDur chan struct{}
}

func main() {
	time2 := inputTime()
	timeRes := time.After(time.Duration(time2) * time.Second)
	timeChannel := make(chan struct{})
	ch := make(chan interface{})
	n := make([]worker, 10)
	for i := 0; i < 10; i++ {
		n[i] = worker{
			Channel: ch,
			timeDur: timeChannel,
		}
		go n[i].work()
	}
	defer close(timeChannel)
	tic := time.Tick(1000 * time.Millisecond)
	for {
		select {
		case <-timeRes:
			timeChannel <- struct{}{}
			close(ch)
			time.Sleep(2 * time.Second)
			return
		case <-tic:
			ch <- rand.Intn(12)
		}
	}
}

func (w worker) work() {
	log.Println("Ready")
	for {
		select {
		case <-w.timeDur:
			log.Println("BB")
			return
		default:
			log.Println(<-w.Channel)
		}
	}
}

func inputTime() int {
	var times int
	fmt.Println("Введите время")
	_, err := fmt.Scanln(&times)
	if err != nil {
		log.Println("fatal error in input!!!!!!!!!")
		return 0
	}
	fmt.Println("Работа программы")
	return times
}
