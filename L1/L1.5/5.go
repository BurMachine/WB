package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type worker struct {
	Channel chan interface{}
}

func main() {
	time2 := inputTime()
	//timeRes := time.After(time.Duration(time2) * time.Second)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time2)*time.Second)
	wg := sync.WaitGroup{}
	ch := make(chan interface{})
	n := make([]worker, 10)
	for i := 0; i < 10; i++ {
		n[i] = worker{
			Channel: ch,
		}
		wg.Add(1)
		go n[i].work(ctx, &wg)
	}
	tic := time.Tick(1000 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			return
		case <-tic:
			ch <- rand.Intn(12)
		}
	}
}

func (w worker) work(ctx context.Context, wg *sync.WaitGroup) {
	log.Println("Ready")
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			if a, ok := <-w.Channel; ok {
				log.Println(a)
			}
			log.Println("BB")
			return
		default:
			if v, ok := <-w.Channel; ok {
				log.Println(v)
			}
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
