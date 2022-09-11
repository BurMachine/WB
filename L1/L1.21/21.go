package main

import "fmt"

// Есть свой тип, которому нужно соответствовать интерфейсу из пакета handlers, чтоб работать с этим пакетом
type myData struct {
	num float64
}

// Метод, который реализует интерфейс
func (md *myData) convertToInt() int {
	return int(md.num)
}

// Создаем адаптер, который должен быть промежуточным звеном между сервисами
type dataAdapter struct {
	myData *myData
}

// Метод адаптера который реализует интерфейс
func (a *dataAdapter) getData() {
	 
	a.myData.convertToInt()
	fmt.Print("Send myData")
}

func main() {
	fmt.Println("aaa")
}
