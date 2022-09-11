package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := 1000000000000, 2000000000000
	num1 := big.NewInt(int64(a))
	num2 := big.NewInt(int64(b))
	sum := new(big.Int)
	sum.Add(num1, num2)
	fmt.Println("a + b = ", sum)
	sub := new(big.Int)
	sub.Sub(num1, num2)
	fmt.Println("a - b = ", sub)
	add := new(big.Int)
	add.Mul(num1, num2)
	fmt.Println("a * b = ", add)
	div := new(big.Int)
	div.Div(num2, num1)
	fmt.Println("b / a = ", div)
}
