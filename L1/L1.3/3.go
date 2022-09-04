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

func (s *arrStruct) pow(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.arr[i] *= s.arr[i]
	s.res += s.arr[i]
	fmt.Println(s.res)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := 0
	//var wg sync.WaitGroup
	var mu sync.Mutex
	s := &arrStruct{
		arr: arr,
		mu:  mu,
		res: res,
	}
	//wg.Add(len(s.arr))
	for i := 0; i < len(s.arr); i++ {
		go s.pow(i)
	}
	//wg.Wait()
	//fmt.Println(s.res)
}
