package main

import "fmt"

func reverse(str string) string {
	runes := []rune(str)                                  // преобразую к слайсу рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // меняю местами все символы
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}
