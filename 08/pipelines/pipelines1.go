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
			// x := <-naturals
			squares <- int(math.Pow(float64(<-naturals), 2))
		}
	}()

	// printer
	for {
		fmt.Println(<-squares)
	}
}
