package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Println("Fibonacci:", fib(10))
	fmt.Println("Greater common divisor:", gcd(10, 8))
}
