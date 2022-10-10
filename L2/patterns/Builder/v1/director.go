package main

/*
	Общий сборщик собирает все, дае если добавить новые типы домов
*/

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetFloorNum()
	return d.builder.GetHouse()
}
