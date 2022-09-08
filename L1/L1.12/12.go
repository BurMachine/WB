package main

import "fmt"

func set(str []string, collection map[string]int) {
	for _, val := range str {
		collection[val] += 1
	}
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	ma := make(map[string]int)
	set(str, ma)
	fmt.Println(ma)
}
