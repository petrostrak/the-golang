package main

import "fmt"

// within the body of the func, the type of
// values is an []int slice.
func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(sum(3))
	fmt.Println(sum(3, 6, 9))
	fmt.Println(sum(3, 6, 9, 12, 15))
	fmt.Println("***")
	a := []int{2, 4, 6, 8, 10}
	fmt.Println(sum(a...))
}
