package main

/*
	Builder - интерфейс строителя, который могут реализовывать множество конкрутных строителей
	В моём случае - это строители домов и иглу
*/
type Builder interface {
	SetDoorType()
	SetFloorNum()
	SetWindowType()
	GetHouse()
}

// Возвращает объект, того типа, который мы создаем(Методы создания реализованы непосредственно для того типа, который создается
func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
