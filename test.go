package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ctx context.Context, ch chan int, num int, exit chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-exit:
			log.Println("Пока")
			return
		default:
			ctx, _ := context.WithTimeout(ctx, 3*time.Second)
			go func(ctx context.Context, ch chan int, num int) {
				select {
				case <-ctx.Done():
					ch <- num
				}
			}(ctx, ch, num)
			time1 := rand.Intn(3)
			time.Sleep(time.Duration(time1) * time.Second)
			fmt.Println("Привет")
		}

	}
}

func main() {
	wg := &sync.WaitGroup{}
	c := context.Background()
	ch := make(chan int, 10)
	exChan := make(chan struct{}, 10)
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Kill, os.Interrupt)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(wg, c, ch, i, exChan)
	}
	go func(ch chan int) {
		for {
			select {
			case a, ok := <-ch:
				if !ok {
					continue
				}
				log.Println("gorutine number:", a)
			}
		}
	}(ch)

	<-signalCh
	for i := 0; i < 10; i++ {
		exChan <- struct{}{}
	}
	defer close(ch)
	wg.Wait()
}
