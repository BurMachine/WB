package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type workers struct {
	Channel     chan interface{}
	ExitChannel chan struct{}
}

func (w *workers) work(id int, wg *sync.WaitGroup) {
	log.Println(id, "Ready")
	defer wg.Done()
	for {
		select {
		case v := <-w.Channel:
			log.Println(v, "worker", id)
		case <-w.ExitChannel:
			if a, ok := <-w.Channel; ok {
				log.Println(a, "worker", id)
			}
			log.Println("exit", id)
			return
		}
	}
}

func main() {
	count := 0
	fmt.Println("Set worker's count:")
	fmt.Scanln(&count)
	fmt.Println("Workers:\n")
	ch := make(chan interface{}, 10)
	exChan := make(chan struct{})
	// os.Signal
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	wg := sync.WaitGroup{}
	// workers
	n := make([]workers, count)
	for i := 0; i < count; i++ {
		n[i] = workers{Channel: ch, ExitChannel: exChan}
		wg.Add(1)
		go n[i].work(i, &wg)
	}
	//for i := 0; i < count; i++ {
	//	go n[i].work(i)
	//}

	tic := time.Tick(1000 * time.Millisecond)
FIRST:
	for {
		select {
		case v := <-signalCh:
			log.Println(v)
			for i := 0; i < count; i++ {
				exChan <- struct{}{}
			}
			close(ch)
			wg.Wait()
			break FIRST
		case <-tic:
			ch <- rand.Intn(100)
		}
	}

	defer close(exChan)
	time.Sleep(1 * time.Second)
}
