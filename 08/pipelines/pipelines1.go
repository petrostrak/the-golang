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
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// // squarer
	go func() {
		for x := range naturals {
			squares <- int(math.Pow(float64(x), 2))
		}
		close(squares)
	}()

	// printer
	for x := range squares {
		fmt.Println(x)
	}
}
