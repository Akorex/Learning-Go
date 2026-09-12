package main

import "fmt"

type MyInt int

func Run14TypeAssertions() {
	var i any
	var mine MyInt = 20

	i = mine

	i2 := i.(MyInt) // checks if the value stored in i can be converted to a MyInt
	fmt.Println(i2 + 1)
	fmt.Printf("%T\n", i2)

	// Safe type assertion with comma-ok idiom
	if i3, ok := i.(string); ok {
		fmt.Println("Is string:", i3)
	} else {
		fmt.Println("Value is not a string")
	}
}
