package VisPkg

type Square struct {
	Side int
}

// Метод из нашей "библиотеки"
func (s *Square) getType() string {
	return "Square"
}

// Ключевой метод, через который будут проходить все изменения
func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}
