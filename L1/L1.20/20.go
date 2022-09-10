package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	array := strings.Split(str, " ")
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	var res string
	for i, val := range array {
		if i == len(array)-1 {
			res += val
			break
		}
		res += val
		res += " "
	}
	return res
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverse(str))
}
