package main

import (
	"fmt"
)

func reverse(s []int) {
	for i := 0; i <= len(s)-1; i++ {
		for j := 0; j < i; j++ {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := []string{"hello"}
	c := []string{"world"}
	// d := []string{"hello", "world"}
	reverse(a[:]) // [:] referes to the whole array
	fmt.Println(a)
	fmt.Println(Equal(b, c))
}
