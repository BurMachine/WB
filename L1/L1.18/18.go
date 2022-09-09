package main

import (
	"fmt"
	"sync"
)

type CountStruct struct {
	sync.RWMutex
	count int
}

func (c *CountStruct) Increment() {
	c.count++
}

func (c *CountStruct) work(n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		go func(i int) {
			c.Lock()
			defer wg.Done()
			defer c.Unlock()
			c.Increment()
			fmt.Println(i, c.count)
		}(i)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	count := &CountStruct{
		RWMutex: sync.RWMutex{},
	}
	wg.Add(100)
	go count.work(100, wg)
	wg.Wait()
	fmt.Println(count.count)
}
