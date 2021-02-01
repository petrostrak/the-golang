package main

import (
	"fmt"
	"math"
)

// Point struct
type Point struct {
	X, Y float64
}

// Path type
type Path []Point

// Distance return the Hopot of two points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance as method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance returns the distance traveled along the path
// We associate methods to any type
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// ScaleBy has a method name (*Point).ScaleBy
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{2, 4}
	q := Point{6, 8}
	fmt.Println(Distance(p, q)) //function call
	fmt.Println(p.Distance(q))  // method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	r := &Point{
		X: 5,
		Y: 5,
	}
	r.ScaleBy(2)
	fmt.Println(*r)
}
