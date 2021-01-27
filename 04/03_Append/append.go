package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is enough room to append
		z = x[:zlen]
	} else {
		// there is not enough room to append
		// so allocate new array by doubling the
		// existing one
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := 6
	fmt.Print(appendInt(a, b))
}
