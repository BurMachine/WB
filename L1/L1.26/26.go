package main

import (
	"fmt"
	"sort"
	"strings"
)

func coincidencesDetector(str string) bool {
	tmp := strings.ToLower(str)
	str1 := []rune(tmp)
	sort.Slice(str1, func(i, j int) bool { return str1[i] < str1[j] }) // Сортировка среза байт из str
	for i := 1; i < len(str1); i++ {
		if str1[i] == str1[i-1] {
			fmt.Println(string(str1))
			return false
		}
	}
	fmt.Println(string(str1))
	return true
}

func main() {
	arr := "cbaСdD"
	coincidencesDetector(arr)
}
