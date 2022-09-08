package main

import "fmt"

type Collection struct {
	ma map[int][]float32
}

func BorderFilling(val float32) (min int) {
	min = int(val/10) * 10
	return
}

func (c *Collection) CollectFilling(arr []float32) {
	for _, val := range arr {
		min := BorderFilling(val)
		c.ma[min] = append(c.ma[min], val)
	}
}

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	ma := make(map[int][]float32)
	coll := &Collection{ma: ma}
	coll.CollectFilling(arr)
	fmt.Println(coll)
}
