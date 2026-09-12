package main

import "fmt"

func Run02ConstantDeclarations() {
	const x int64 = 10

	const (
		idKey   = "id"
		nameKey = "name"
	)

	const z = 20 * 10
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(idKey, nameKey, z)

	// In Go, constants cannot be reassigned (uncommenting causes compile errors):
	// x = x + 1
	// y = "bye"
}
