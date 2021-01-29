package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle: Circle{Point: Point{X: 5, Y: 10}, Radius: 6}, Spokes: 20}
	w = Wheel{Circle{Point{2, 4}, 5}, 18}

	w.X = 11
	fmt.Printf("%#v\n", w) // Wheel{Circle:main.Circle{Point:main.Point{X:11, Y:4}, Radius:5}, Spokes:18}
}
