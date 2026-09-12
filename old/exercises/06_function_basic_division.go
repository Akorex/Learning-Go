package main

import "fmt"

func divBasic(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func Run06FunctionBasicDivision() {
	result := divBasic(5, 2)
	fmt.Println(result)
}
