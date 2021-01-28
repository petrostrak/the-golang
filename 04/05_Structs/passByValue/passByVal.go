package main

import "fmt"

type person struct {
	fName, lName string
}

// changeNameV takes a struct and changes the first name.
// The struct is passed by value, meaning it creates a copy
// of the p struct
func changeNameV(p person) {
	p.fName = "Petros"
}

// changeNameP thats a pointer to a struct and changes the
// first name.
func changeNameP(p *person) {
	p.fName = "Petros"
}

func main() {
	p := person{
		fName: "Pit",
		lName: "Trak",
	}

	changeNameV(p)
	fmt.Println(p)
	changeNameP(&p)
	fmt.Println(p)
}
