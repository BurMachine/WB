package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func workerOne(ch1, ch2 chan int, exit chan struct{}, wg *sync.WaitGroup) {
	log.Println("First worker start")
	defer wg.Done()
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case v, ok := <-ch1:
			if !ok {
				continue
			}
			log.Println("Worker one: ", v)
			ch2 <- v * v
		case <-exit:
			v, ok := <-ch1
			if ok {
				log.Println("Worker one exit work: ", v)
				ch2 <- v * v
			}
			close(ch2)
			log.Println("worker one exit")
			return
		}
	}
}

func workerTwo(ch2 chan int, exit chan struct{}, wg *sync.WaitGroup) {
	log.Println("First worker start")
	defer wg.Done()
	for {
		time.Sleep(100 * time.Millisecond)

		select {
		case v, ok := <-ch2:
			if !ok {
				continue
			}
			log.Println("Worker two: ", v)
		case <-exit:
			v, ok := <-ch2
			if ok {
				log.Println("Worker two exit work: ", v)
			}
			log.Println("worker two exit")
			return
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	exChan := make(chan struct{})
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Kill, os.Interrupt)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go workerOne(ch1, ch2, exChan, wg)
	go workerTwo(ch2, exChan, wg)

	go func(chan int, chan struct{}) {
		for {
			select {
			case <-exChan:
				close(ch1)
				return
			default:
				time.Sleep(50 * time.Millisecond)
				ch1 <- rand.Intn(10)
			}
		}
	}(ch1, exChan)

	<-signalCh

	for i := 0; i < 3; i++ {
		exChan <- struct{}{}
	}
	wg.Wait()
	log.Println("EXIT")
}
