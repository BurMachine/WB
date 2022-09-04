package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// композиция структур
type Action struct {
	Human
	Act string
}

func (h *Human) GetInfo() {
	fmt.Println(h.Age, "years old\n", h.Name)
}
func (h *Human) SetInfo(name string, age int) {
	h.Name = name
	h.Age = age
}
func main() {
	// структура action имеет методы human
	actor := &Action{
		Human: Human{
			Name: "Artur",
			Age:  20,
		},
		Act: "build",
	}
	actor.GetInfo()
	actor.SetInfo("ivan", 33)
	actor.GetInfo()
}
