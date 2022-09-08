package main

import "fmt"

func intersection(set1, set2 []int) []int {
	ma1 := make(map[int]int)
	ma2 := make(map[int]int)
	res := make([]int, 0)
	for _, val := range set1 {
		ma1[val] += 1
	}
	for _, val := range set2 {
		ma2[val] += 1
		if ma1[val] > 0 && ma2[val] > 0 {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	set1 := []int{1, 1, 1, 1, 1, 1, 9, 2, 3}
	set2 := []int{2, 2, 4, 5, 6, 7, 8, 9, 1}
	res := intersection(set1, set2)
	fmt.Println(res)
}
