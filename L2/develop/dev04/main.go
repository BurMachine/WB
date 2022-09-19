package main

import (
	"burmachine/anagram/anagramm"
	"fmt"
)

func main() {
	str := []string{"кот", "ток", "окт", "кит", "тик", "кто", "сок", "Сок"}
	ma := anagramm.FindAnagram(&str)
	fmt.Println(ma)
}
