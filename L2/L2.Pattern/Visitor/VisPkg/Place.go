package VisPkg

type Place interface {
	Accept(visitor Visitor) string
}

type Sushi struct {
	Name string
}

type Pizza struct {
	Name string
}

type Burger struct {
	Name string
}

func (s *Sushi) BuySushi() string {
	return "sushiiiiiiiuiuiiiiiuiuiiiiiuiui!!!!!!!!!!!!!!!!!!!"
}

func (p *Pizza) BuyPizza() string {
	return "pizzzzzaaaaaa!!!!!!!!!!!!!!!!!!!)))))·)()(((()())))))!!!!!"
}
func (b *Burger) BuyBurger() string {
	return "BEEEERGEER!!!!!!!!!!!!!!!!!!!)))))·)()(((()())))))!!!!!"
}

func (s *Sushi) Accept(visitor Visitor) string {
	return visitor.VisitSushiBar(s)
}

func (p *Pizza) Accept(visitor Visitor) string {
	return visitor.VisitPizzeria(p)
}

func (b *Burger) Accept(visitor Visitor) string {
	return visitor.VisitBurgerBar(b)
}

//************************************************************************************************************//
//                                   				CITY                                                      //

type City struct {
	places []Place
}

func (c City) Add(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

//                                                                                                            //
//************************************************************************************************************//
