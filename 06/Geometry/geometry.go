package main

import (
	"fmt"
	"math"
)

// Point struct
type Point struct {
	X, Y float64
}

// Distance return the Hopot of two points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance as method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{2, 4}
	q := Point{6, 8}
	fmt.Println(Distance(p, q)) //function call
	fmt.Println(p.Distance(q))  // method call
}
