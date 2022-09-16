package main

import "simple/VisPkg/VisPkg"

func main() {
	square := &VisPkg.Square{Side: 2}
	circle := &VisPkg.Circle{Radius: 4}
	triangle := &VisPkg.Triangle{
		L: 2,
		B: 2,
	}

	areaCalculator := &VisPkg.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	triangle.Accept(areaCalculator)

	middleCoordinates := &VisPkg.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	triangle.Accept(middleCoordinates)
}
