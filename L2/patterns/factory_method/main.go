package main

import "fmt"

//************************* INTERFACE

type IComputer interface {
	getMemory() int
	getName() string
	setMemory(int)
	setName(string)
}

//************************* Обобщенный продукт

type Computer struct {
	Memory int
	Name   string
}

func (c Computer) getMemory() int {
	return c.Memory
}

func (c Computer) getName() string {
	return c.Name
}

func (c *Computer) setMemory(i int) {
	c.Memory = i
}

func (c *Computer) setName(s string) {
	c.Name = s
}

//************************* Конкретные продукты

type Intel struct {
	Computer
}

func NewIntel() IComputer {
	return &Intel{Computer{
		Memory: 1024,
		Name:   "Intel computer",
	}}
}

type Acer struct {
	Computer
}

func NewAcer() IComputer {
	return &Acer{Computer{
		Memory: 2048,
		Name:   "Acer Nitro 5",
	}}
}

//************************* Фабрика

func GetComputer(s string) IComputer {
	if s == "Intel" {
		return NewIntel()
	} else if s == "Acer" {
		return NewAcer()
	}
	return nil
}

//************************* Клиентский код

func main() {
	intel := GetComputer("Intel")
	acer := GetComputer("Acer")

	Printer(intel)
	Printer(acer)
}

func Printer(computer IComputer) {
	fmt.Println(computer.getName(), computer.getMemory())
}
