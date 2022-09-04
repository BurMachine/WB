package main

import (
	"fmt"
	"sync"
)

type arrStruct struct {
	arr []int
	res int
	mu  sync.Mutex
}

func (s *arrStruct) pow(i int, wg *sync.WaitGroup) {
	// закрываю и откладываю открытие мьютекса
	s.mu.Lock()
	defer s.mu.Unlock()
	defer wg.Done() // уменьшаю waitgroup
	s.arr[i] *= s.arr[i]
	s.res += s.arr[i]
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	s := &arrStruct{
		arr: arr,
		mu:  mu,
		res: res,
	}
	a := len(s.arr)
	wg.Add(a)
	// разделяю на горутины
	for i := 0; i < len(s.arr); i++ {
		go s.pow(i, &wg)
	}
	wg.Wait() // жду waitgroup(main)
	fmt.Println(s.res)
}
