package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func cToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100C
	boilingF := cToF(BoilingC)
	fmt.Printf("%g\n", boilingF-cToF(FreezingC)) // 180F
	// fmt.Printf("%g'n", boilingF-FreezingC) // compile error type mismatch

	// even though c and f have the same underlying type (float64)
	// they are not the same type
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // compile error type mismatch
	fmt.Println(c == Celsius(f)) // true
}
