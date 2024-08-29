package main

import "fmt"

type Inner struct {
	A int
}

type Outer struct {
	Inner
	S string
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}

	fmt.Println(o.Double())

}
