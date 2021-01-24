// echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for i, args := range os.Args[1:] {
		s += sep + args
		sep = " "
		fmt.Println(i)
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	fmt.Println(os.Args[0:])
}

func main() {
	echo2()
}
