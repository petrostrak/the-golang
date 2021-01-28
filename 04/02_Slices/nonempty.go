package main

import "fmt"

func nonEmpty(str []string) []string {
	i := 0
	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

func nonEmptyAlt(str []string) []string {
	out := str[:0] // creating an zero-length slice of original
	for _, s := range str {
		if s != "" {
			out = append(out, s) // push s
		}
	}
	return out
}

// removeElem removes an element from the middle of the
// given slice, preserving the order of the remaining
// elements
func removeElem(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(removeElem(a, 3))
}
