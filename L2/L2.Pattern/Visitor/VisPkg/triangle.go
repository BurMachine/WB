package VisPkg

type Triangle struct {
	L int
	B int
}

// Ключевой метод, через который будут проходить все изменения
func (t *Triangle) Accept(v Visitor) {
	v.visitForTriangle(t)
}

func (t *Triangle) getType() string {
	return "Triangle"
}
