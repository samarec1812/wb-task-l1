package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/


type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(one, two *Point) float64 {
	a := math.Pow(two.x - one.x, 2)
	b := math.Pow(two.y - one.y, 2)
	return math.Sqrt(a + b)
}

func main() {
	pointOne := NewPoint(2, 3)
	pointTwo := NewPoint(4, 5)
	fmt.Println("Distance: ", Distance(pointOne, pointTwo))
}