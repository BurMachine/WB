package main

var justString string

func createHugeString(len int) string {
	var str string
	for i := 0; i < len; i++ {
		str += "q"
	}
	// some code
	return str
}

// У нас будет 924 неисполользуемых байта памяти, которые не сможет очистить сборщик мусора
// Они будут висеть в столько, сколько работает наша программа
// Поэтому лучше сразу выделять столько памяти, сколько нужно, или вместо среза копировать 100 элементов
// Также тут может возникнуть проблема с юникодом т.к. он занимает 2 байта в памяти
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

// OR
