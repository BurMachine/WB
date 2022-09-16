package factory

/*
	Интерфейс продукта
*/

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
