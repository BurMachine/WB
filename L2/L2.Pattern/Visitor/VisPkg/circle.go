package VisPkg

type Circle struct {
	Radius int
}

// Ключевой метод, через который будут проходить все изменения
func (c *Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
