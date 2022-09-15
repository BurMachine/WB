package VisPkg

type Visitor interface {
	VisitSushiBar(p *Sushi) string
	VisitPizzeria(p *Pizza) string
	VisitBurgerBar(p *Burger) string
}
