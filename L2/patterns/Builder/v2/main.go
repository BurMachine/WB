package main

import "fmt"

//************************** ИНТЕРФЕЙС СТРОИТЕЛЯ

type Builder interface {
	SetHp()
	SetTwistingMoment()

	GetAuto() Auto
}

func getAuto(name string) Builder {
	if name == "supra" {
		return newSupra()
	} else if name == "challenger" {
		return newChallenger()
	}
	return nil
}

//*************************** ОБЪЕКТ КОТОРЫЙ СТРОЯТ

type Auto struct {
	HorsePower     int
	TwistingMoment int
	Name           string
}

//*************************** Конкретный строитель Toyota

type Toyota struct {
	HP          int
	TwistMoment int
	Name        string
}

func (t *Toyota) SetHp() {
	t.HP = 399
}

func (t *Toyota) SetTwistingMoment() {
	t.TwistMoment = 510
}

func (t *Toyota) GetAuto() Auto {
	return Auto{
		HorsePower:     t.HP,
		TwistingMoment: t.TwistMoment,
		Name:           t.Name,
	}
}

func newSupra() *Toyota {
	return &Toyota{Name: "Toyota Supra 80 DTM"}
}

//*************************** Конкрутный строитель Dodge

type Dodge struct {
	HP          int
	TwistMoment int
	Name        string
}

func (d *Dodge) SetHp() {
	d.HP = 600
}

func (d *Dodge) SetTwistingMoment() {
	d.TwistMoment = 800
}

func (d *Dodge) GetAuto() Auto {
	return Auto{
		HorsePower:     d.HP,
		TwistingMoment: d.TwistMoment,
		Name:           d.Name,
	}
}

func newChallenger() *Dodge {
	return &Dodge{Name: "Dodge Challenger Hellcat"}
}

//*************************** Общий сборщик Director

type Director struct {
	Builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{Builder: b}
}

func (d *Director) SetBuilderToDirector(b Builder) {
	d.Builder = b
}

func (d *Director) BuildMachine() Auto {
	d.Builder.SetHp()
	d.Builder.SetTwistingMoment()
	return d.Builder.GetAuto()
}

//*************************** Клиентский код

func main() {
	supra := getAuto("supra")
	challenger := getAuto("challenger")

	builder := newDirector(supra)
	machine := builder.BuildMachine()

	fmt.Println(machine.Name, machine.HorsePower, machine.TwistingMoment)

	builder.SetBuilderToDirector(challenger)
	machine = builder.BuildMachine()

	fmt.Println(machine.Name, machine.HorsePower, machine.TwistingMoment)
}
