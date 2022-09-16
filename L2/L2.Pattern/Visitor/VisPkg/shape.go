package VisPkg

/*
	Shape
	Интерфейс, который реализует каждая наша фигура
*/
type Shape interface {
	getType() string
	Accept(visitor Visitor)
}
