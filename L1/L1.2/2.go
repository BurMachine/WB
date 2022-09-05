package main

import (
	"fmt"
	"sync"
)

type arrStruct struct {
	arr []int
	mu  sync.Mutex
}

func (s *arrStruct) pow(i int, wg *sync.WaitGroup) {
	s.mu.Lock()
	defer wg.Done()
	defer s.mu.Unlock()
	s.arr[i] *= s.arr[i]
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mu sync.Mutex
	s := &arrStruct{
		arr: arr,
		mu:  mu,
	}
	for i := 0; i < len(s.arr); i++ {
		wg.Add(1)
		go s.pow(i, &wg)
	}
	wg.Wait()
	fmt.Println(s.arr)
}
