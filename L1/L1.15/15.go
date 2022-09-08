package main

import "fmt"

var justString string

func createHugeString(int) string {
	var str string
	fmt.Println(str)
	// some code
	return str
}

// У нас будет 924 неисполользуемых байта памяти, которые не сможет очистить сборщик мусора
// Они будут висеть в столько, сколько работает наша программа
// Поэтому лучше сразу выделять столько памяти, сколько нужно, или вместо среза копировать 100 элементов

// TASK REALIZATION
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// CORRECT VERSION

func someFunc() {
	justString = createHugeString(100)
}

func main() {
	someFunc()
}
