package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

// Конструктор
func ConstructPoint() *Point {
	return &Point{
		x: 0,
		y: 0,
	}
}

// Сеттер
func (p *Point) SetPoint(x, y int) {
	p.x = x
	p.y = y
}

// Гетеры
func (p Point) GetPointX() int {
	return p.x
}

func (p Point) GetPointY() int {
	return p.y
}

// Заполнение min/max
func fillingMinMax(point1, point2 Point) (int, int, int, int) {
	var minX, maxX, minY, maxY int
	if point1.GetPointX() > point2.GetPointX() {
		minX = point2.GetPointX()
		maxX = point1.GetPointX()
	} else {
		minX = point1.GetPointX()
		maxX = point2.GetPointX()
	}
	if point1.GetPointY() > point2.GetPointY() {
		minY = point2.GetPointY()
		maxY = point1.GetPointY()
	} else {
		minY = point1.GetPointY()
		maxY = point2.GetPointY()
	}
	return maxX, minX, maxY, minY
}

// Расчет расстояния
func CalculateDistance(point1, point2 Point) float64 {
	maxX, minX, maxY, minY := fillingMinMax(point1, point2)
	distance := math.Sqrt(float64((maxX-minX)*(maxX-minX) + (maxY-minY)*(maxY-minY)))
	return distance
}

func main() {
	a := ConstructPoint()
	b := ConstructPoint()
	a.SetPoint(1, 2)
	b.SetPoint(3, 4)
	dist := CalculateDistance(*a, *b)
	fmt.Println(dist)
}
