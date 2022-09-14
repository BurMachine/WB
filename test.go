package main

import "fmt"

type demo struct {
	s string
}

func change(a []demo) {
	for _, d := range a {
		d.s = "246"
	}
}

func changeP(a []*demo) {
	for _, d := range a {
		d.s = "246"
	}
}

func main() {
	d := make([]demo, 5)
	dP := make([]*demo, 5)
	for i := 0; i < 5; i++ {
		d[i].s = "123"
		dP[i] = new(demo)
		dP[i].s = "123"
	}
	change(d)
	changeP(dP)
	fmt.Println(d)
	for _, d2 := range dP {
		fmt.Println((*d2).s)
	}
}
