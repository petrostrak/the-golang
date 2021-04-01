package main

import (
	"fmt"
	"math"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// // squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			// x := <-naturals
			squares <- int(math.Pow(float64(x), 2))
		}
		close(squares)
	}()

	// printer
	for {
		fmt.Println(<-squares)
	}
}
