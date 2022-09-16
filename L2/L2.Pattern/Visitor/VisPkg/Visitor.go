package VisPkg

/*
	Visitor
	Функции visit помогают нам добавлять функционал для наших типов
*/
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Triangle)
}
