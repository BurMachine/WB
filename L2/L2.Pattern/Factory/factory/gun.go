package factory

/*
	Конкретный продукт в общем виде
*/

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) GetPower() int {
	return g.Power
}
