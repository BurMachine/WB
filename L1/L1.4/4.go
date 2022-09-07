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
		time.Sleep(100 * time.Millisecond)
		select {
		case v, ok := <-w.Channel:
			if !ok {
				continue
			}
			log.Println(v, "worker", id)
		case <-w.ExitChannel:
			for i := range w.Channel {
				log.Println(i, "worker", id)
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
	ch := make(chan interface{}, 100)
	exChan := make(chan struct{})
	// os.Signal
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	wg := &sync.WaitGroup{}
	// workers
	n := make([]workers, count)
	for i := 0; i < count; i++ {
		n[i] = workers{Channel: ch, ExitChannel: exChan}
		wg.Add(1)
		go n[i].work(i, wg)
	}

	go func(ch chan interface{}) {
		tic := time.Tick(10 * time.Millisecond)
		for {
			select {
			case <-exChan:
				close(ch)
				return
			default:
				<-tic
				ch <- rand.Intn(100)
			}
		}
	}(ch)

	<-signalCh
	for i := 0; i < count+1; i++ {
		exChan <- struct{}{}
	}
	wg.Wait()

	defer close(exChan)
}
