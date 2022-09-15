package VisPkg

type People struct {
	Name string
}

func (p2 People) VisitSushiBar(p *Sushi) string {
	return p.BuySushi()
}

func (p2 People) VisitPizzeria(p *Pizza) string {
	return p.BuyPizza()
}

func (p2 People) VisitBurgerBar(p *Burger) string {
	return p.BuyBurger()
}
