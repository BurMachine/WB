package main

import (
	"fmt"
)

func getIntersection(a []string, b []string, mode byte) []string {
	m := make(map[string]byte)

	for _, k := range a {
		m[k] += 1
	}

	for _, k := range b {
		m[k] += 2
	}

	result := []string{}

	if mode == 4 {
		for k, v := range m {
			if v < 3 {
				result = append(result, k)
			}
		}
	} else {
		for k, v := range m {
			if v == mode {
				result = append(result, k)
			}
		}
	}

	return result
}

func main() {

	a := []string{"a", "b", "c", "d"}
	b := []string{"c", "d", "e", "f"}

	fmt.Println("left", getIntersection(a, b, 1))
	fmt.Println("right", getIntersection(a, b, 2))
	fmt.Println("left & right", getIntersection(a, b, 3))
	fmt.Println("left ^ right", getIntersection(a, b, 4))
}
