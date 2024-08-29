package main

import "fmt"

func makePointer[T any](t T) *T {
	return &t
}

func main() {
	var x int = 10
	var y bool = true
	var z int = 15

	pointerX := &x
	pointerY := &y
	pointerZ := &z

	fmt.Println(pointerX)
	fmt.Println(pointerY)

	fmt.Println(*pointerX + *pointerZ)

	type person struct {
		firstName  string
		middleName *string
		lastName   string
	}

	p := person{
		firstName:  "Akorede",
		middleName: makePointer("Idris"),
		lastName:   "Adewole",
	}

}
