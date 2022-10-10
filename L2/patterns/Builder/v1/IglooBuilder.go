package main

// Тут всё аналогично NormalBuilder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) SetFloorNum() {
	b.floor = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
