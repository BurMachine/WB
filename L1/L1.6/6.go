package main

import (
	"log"
	"time"
)

//*****************************************************************//
// First way - closing channel

//func FirstGo(ch chan int) {
//	n := 0
//	for {
//		select {
//		case ch <- n:
//			n++
//		case <-ch:
//			return
//		}
//	}
//}
//
//func main() {
//	ch := make(chan int)
//	go FirstGo(ch)
//	log.Println(<-ch)
//	log.Println(<-ch)
//	log.Println(<-ch)
//	close(ch)
//}

//*****************************************************************//
// Second way - context(any)

//func SecondGo(ctx context.Context) {
//	for {
//		time.Sleep(100 * time.Millisecond)
//		select {
//		case <-ctx.Done():
//			return
//		default:
//			fmt.Println(rand.Intn(12))
//		}
//	}
//}
//
//func main() {
//	ctx, close := context.WithCancel(context.Background())
//	go SecondGo(ctx)
//	time.Sleep(1 * time.Second)
//	close()
//}

//*****************************************************************//
// Third way - main gorutine

func ThirdGo() {
	for {
		time.Sleep(100 * time.Millisecond)
		log.Println("GO")
	}
}

func main() {
	go ThirdGo()
	time.Sleep(1 * time.Second)
}

//*****************************************************************//
