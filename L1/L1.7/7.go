package main

import (
	"log"
	"math/rand"
	"sync"
)

type mapInfo struct {
	*sync.RWMutex
	ma map[int]int
}

func (m *mapInfo) Recording(key, value int, wg *sync.WaitGroup) {
	m.Lock()
	defer wg.Done()
	defer m.Unlock()
	m.ma[key] = value
}

func main() {
	mapa := make(map[int]int)
	wg := &sync.WaitGroup{}
	m := &mapInfo{
		RWMutex: &sync.RWMutex{},
		ma:      mapa,
	}
	for i := 0; i < 100; i++ {
		j := rand.Intn(100)
		wg.Add(1)
		go m.Recording(i, j, wg)
	}
	wg.Wait()
	log.Println(m.ma)
}
