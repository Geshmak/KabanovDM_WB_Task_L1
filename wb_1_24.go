package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(2, 7)

	fmt.Println(point1.x, "  ", point1.y)
	fmt.Println(point2.x, "  ", point2.y)
	fmt.Println(Distance(point1, point2))
}
