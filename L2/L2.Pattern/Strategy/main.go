package main

import (
	"fmt"
	"simple/Strategy/strategy"
)

/*
	main
	 Мы определяем поведения метода Operate() по ходу выполнения программы
*/
func main() {
	add := strategy.Operation{strategy.Addiction{}}
	fmt.Println(add.Operate(3, 5))
	mul := strategy.Operation{strategy.Multiplication{}}
	fmt.Println(mul.Operate(3, 5))
}
