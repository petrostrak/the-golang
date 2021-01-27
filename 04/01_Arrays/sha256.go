package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum224([]byte("x"))
	c2 := sha256.Sum224([]byte("X"))
	// The two inputs differ by only a single bit
	// but approx half the bits are different in the
	// digest
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}
