package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20

	i = mine

	i2 := i.(MyInt) // checks if the value stored in i can be converted to a MyInt
	fmt.Println(i2 + 1)
	fmt.Printf("%T ", i2)
}
