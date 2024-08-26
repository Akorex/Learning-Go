package main

import "fmt"

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}

	return num / denom
}

func main() {
	result := div(5, 2)

	fmt.Println(result)
}
