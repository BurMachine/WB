package VisPkg

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(square *Square) {
	a.area = square.Side * square.Side
	fmt.Printf("Area of sqare with side: %d = %d\n", square.Side, a.area)
}

func (a *AreaCalculator) visitForCircle(circle *Circle) {
	rad := float64(circle.Radius)
	a.area = int(math.Pi * rad * rad)
	fmt.Printf("Rounded area of %d cm radius circle = %d\n", circle.Radius, a.area)
}

func (a *AreaCalculator) visitForTriangle(triangle *Triangle) {
	a.area = int(triangle.B * triangle.L / 2)
	fmt.Printf("Area of orthogonal triangle = %d\n", a.area)
}
