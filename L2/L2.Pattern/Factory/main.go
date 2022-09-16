package main

import (
	"fmt"
	"simple/factory/factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")
	printDetails(ak47)
	fmt.Println()
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Println("Gun:", g.GetName())
	fmt.Println("Power:", g.GetPower())
}
