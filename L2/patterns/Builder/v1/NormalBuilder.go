package main

/*
	NormalBuilder - конкретный строитель(строитель нормальных домов)
*/
type NormalBuilder struct {
	windowType string
	doorType   string
	floorNum   int
}

// Конструктор
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden door"
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden door"
}

func (b *NormalBuilder) SetFloorNum() {
	b.floorNum = 2
}

// GetHouse возвращает заполненный объект класса house
func (b *NormalBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floorNum,
	}
}
