package main

import (
	"fmt"
	"sync"
	"time"
)

type arrStruct struct {
	arr []int
	mu  sync.Mutex
}

//func pow(number int) {
//	fmt.Printf("%d\n", number*number)
//}

func (s *arrStruct) pow(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.arr[i] *= s.arr[i]
	//fmt.Printf("%d\n", s.arr)
	//s.mu.Unlock()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var mu sync.Mutex
	s := &arrStruct{
		arr: arr,
		mu:  mu,
	}
	for i := 0; i < len(s.arr); i++ {
		go s.pow(i)
	}
	fmt.Println(s.arr)

	time.Sleep(1 * time.Second)
}
